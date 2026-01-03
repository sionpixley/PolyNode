package utilities

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sionpixley/PolyNode/internal/constants/command"
	"github.com/sionpixley/PolyNode/internal/models"
)

func ConvertToCommand(commandStr string) models.Command {
	switch commandStr {
	case "add":
		return command.Add
	case "current":
		return command.Current
	case "install":
		return command.Install
	case "ls":
		fallthrough
	case "list":
		return command.List
	case "rm":
		fallthrough
	case "remove":
		return command.Remove
	case "search":
		return command.Search
	case "temp":
		return command.Temp
	case "use":
		return command.Use
	default:
		return command.Other
	}
}

func ConvertToSemanticVersion(version string) string {
	if version[0] == 'v' {
		return version
	}

	return "v" + version
}

func ExtractFile(source string, destination string) error {
	err := os.RemoveAll(destination)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

	if strings.HasSuffix(source, ".gz") {
		err = ExtractGzip(source, destination)
		if err != nil {
			return err
		}
	} else {
		err = ExtractZip(source, destination)
		if err != nil {
			return err
		}
	}

	return os.RemoveAll(source)
}

func ExtractGzip(source string, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		target := filepath.Join(destination, stripTopDir(header.Name))

		switch header.Typeflag {
		case tar.TypeDir:
			if e := os.MkdirAll(target, os.FileMode(header.Mode)); e != nil {
				return e
			}
		case tar.TypeReg:
			if e := os.MkdirAll(filepath.Dir(target), os.FileMode(header.Mode)); e != nil {
				return e
			}
			outFile, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, e2 := io.Copy(outFile, tarReader); e2 != nil {
				_ = outFile.Close()
				return e2
			}
			_ = outFile.Close()
		case tar.TypeSymlink:
			if e := os.MkdirAll(filepath.Dir(target), os.FileMode(header.Mode)); e != nil {
				return e
			}
			if e2 := os.Symlink(header.Linkname, target); e2 != nil {
				return e2
			}
		case tar.TypeLink:
			if e := os.MkdirAll(filepath.Dir(target), os.FileMode(header.Mode)); e != nil {
				return e
			}
			if e2 := os.Link(header.Linkname, target); e2 != nil {
				return e2
			}
		default:
			// Do nothing.
		}
	}

	return nil
}

func ExtractZip(source string, destination string) error {
	zipReader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		target := filepath.Join(destination, stripTopDir(strings.ReplaceAll(file.Name, "\\", "/")))

		if file.FileInfo().IsDir() {
			if e := os.MkdirAll(target, file.Mode()); e != nil {
				return e
			}
		} else if file.Mode()&os.ModeSymlink != 0 {
			if e := os.MkdirAll(filepath.Dir(target), file.Mode()); e != nil {
				return e
			}

			src, err := file.Open()
			if err != nil {
				return err
			}

			link, err := io.ReadAll(src)
			if err != nil {
				_ = src.Close()
				return err
			}

			if e2 := os.Symlink(string(link), target); e2 != nil {
				_ = src.Close()
				return e2
			}

			_ = src.Close()
		} else {
			if e := os.MkdirAll(filepath.Dir(target), file.Mode()); e != nil {
				return e
			}

			src, err := file.Open()
			if err != nil {
				return err
			}

			dist, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
			if err != nil {
				_ = src.Close()
				return err
			}

			if _, e2 := io.Copy(dist, src); e2 != nil {
				_ = src.Close()
				_ = dist.Close()
				return e2
			}

			_ = src.Close()
			_ = dist.Close()
		}
	}

	return nil
}

func KnownCommand(comm string) bool {
	return ConvertToCommand(comm) != command.Other
}

func ValidVersionFormat(version string) bool {
	if version[0] == 'v' {
		version = version[1:]
	}

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return false
	}

	validChars := map[rune]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
	}
	for _, part := range parts {
		if part == "" {
			return false
		}
		for _, char := range part {
			if _, exists := validChars[char]; !exists {
				return false
			}
		}
	}

	return true
}

func stripTopDir(path string) string {
	parts := strings.SplitN(path, "/", 2)
	if len(parts) == 2 {
		return parts[1]
	}

	return path
}

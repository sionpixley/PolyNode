package utilities

import (
	"archive/tar"
	"compress/gzip"
	"github.com/sionpixley/PolyNode/internal"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
)

func ConvertToCommand(commandStr string) models.Command {
	switch commandStr {
	case "add":
		return constants.Add
	case "current":
		return constants.Current
	case "install":
		return constants.Install
	case "ls":
		fallthrough
	case "list":
		return constants.List
	case "rm":
		fallthrough
	case "remove":
		return constants.Remove
	case "search":
		return constants.Search
	case "temp":
		return constants.Temp
	case "use":
		return constants.Use
	default:
		return constants.NAComm
	}
}

func ConvertToSemanticVersion(version string) string {
	if version[0] == 'v' {
		return version
	} else {
		return "v" + version
	}
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

	err = exec.Command("tar", "-xf", source, "-C", destination, "--strip-components=1").Run()
	if err != nil {
		return err
	}

	return os.RemoveAll(source)
}

func ExtractGzip(source string, destination string) error {
	err := os.RemoveAll(destination)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

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
				outFile.Close()
				return e2
			}
			outFile.Close()
		default:
			// Do nothing.
		}
	}

	return nil
}

func IsKnownCommand(command string) bool {
	return ConvertToCommand(command) != constants.NAComm
}

func IsValidVersionFormat(version string) bool {
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
	parts := strings.SplitN(path, internal.PathSeparator, 2)
	if len(parts) == 2 {
		return parts[1]
	} else {
		return path
	}
}

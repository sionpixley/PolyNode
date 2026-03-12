package node

import (
	"slices"
	"testing"

	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
)

func TestConvertOSAndArchToNodeVersionFile_AIX(t *testing.T) {
	expected := "aix-ppc64"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.AIX, arch.PPC64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Linux_ARM64(t *testing.T) {
	expected := "linux-arm64"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Linux, arch.ARM64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Linux_PPC64LE(t *testing.T) {
	expected := "linux-ppc64le"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Linux, arch.PPC64LE)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Linux_S390X(t *testing.T) {
	expected := "linux-s390x"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Linux, arch.S390X)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Linux_X64(t *testing.T) {
	expected := "linux-x64"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Linux, arch.X64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Mac_ARM64(t *testing.T) {
	expected := "osx-arm64-tar"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Mac, arch.ARM64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Mac_X64(t *testing.T) {
	expected := "osx-x64-tar"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Mac, arch.X64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_UnsupportedArch(t *testing.T) {
	_, err := convertOSAndArchToNodeVersionFile(opsys.Linux, arch.PPC64)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestConvertOSAndArchToNodeVersionFile_UnsupportedOS(t *testing.T) {
	_, err := convertOSAndArchToNodeVersionFile(opsys.Other, arch.PPC64)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestConvertOSAndArchToNodeVersionFile_Windows_ARM64(t *testing.T) {
	expected := "win-arm64-zip"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Windows, arch.ARM64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertOSAndArchToNodeVersionFile_Windows_X64(t *testing.T) {
	expected := "win-x64-zip"
	actual, err := convertOSAndArchToNodeVersionFile(opsys.Windows, arch.X64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetAllNodeVersionsForOSAndArch(t *testing.T) {
	httpWrapper := new(models.HTTPMock)
	osWrapper := new(models.OSMockNotExist)
	config := models.NewPolyNodeConfig(osWrapper)

	expected := []models.NodeVersion{
		{
			Version: "v25.8.1",
			Files:   []string{"aix-ppc64", "headers", "linux-arm64", "linux-ppc64le", "linux-s390x", "linux-x64", "osx-arm64-tar", "osx-x64-pkg", "osx-x64-tar", "src", "win-arm64-7z", "win-arm64-zip", "win-x64-7z", "win-x64-exe", "win-x64-msi", "win-x64-zip"},
			LTS:     false,
		},
		{
			Version: "v25.8.0",
			Files:   []string{"aix-ppc64", "headers", "linux-arm64", "linux-ppc64le", "linux-s390x", "linux-x64", "osx-arm64-tar", "osx-x64-pkg", "osx-x64-tar", "src", "win-arm64-7z", "win-arm64-zip", "win-x64-7z", "win-x64-exe", "win-x64-msi", "win-x64-zip"},
			LTS:     false,
		},
		{
			Version: "v24.14.0",
			Files:   []string{"aix-ppc64", "headers", "linux-arm64", "linux-ppc64le", "linux-s390x", "linux-x64", "osx-arm64-tar", "osx-x64-pkg", "osx-x64-tar", "src", "win-arm64-7z", "win-arm64-zip", "win-x64-7z", "win-x64-exe", "win-x64-msi", "win-x64-zip"},
			LTS:     true,
		},
	}
	actual, err := getAllNodeVersionsForOSAndArch(opsys.Linux, arch.ARM64, config, httpWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	if !slicesEqual(actual, expected) {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestGetArchiveName_AIX(t *testing.T) {
	expected := "aix-ppc64.tar.gz"
	actual, err := getArchiveName(opsys.AIX, arch.PPC64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Linux_ARM64(t *testing.T) {
	expected := "linux-arm64.tar.gz"
	actual, err := getArchiveName(opsys.Linux, arch.ARM64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Linux_PPC64LE(t *testing.T) {
	expected := "linux-ppc64le.tar.gz"
	actual, err := getArchiveName(opsys.Linux, arch.PPC64LE)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Linux_S390X(t *testing.T) {
	expected := "linux-s390x.tar.gz"
	actual, err := getArchiveName(opsys.Linux, arch.S390X)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Linux_X64(t *testing.T) {
	expected := "linux-x64.tar.gz"
	actual, err := getArchiveName(opsys.Linux, arch.X64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Mac_ARM64(t *testing.T) {
	expected := "darwin-arm64.tar.gz"
	actual, err := getArchiveName(opsys.Mac, arch.ARM64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_Mac_X64(t *testing.T) {
	expected := "darwin-x64.tar.gz"
	actual, err := getArchiveName(opsys.Mac, arch.X64)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetArchiveName_UnsupportedArch(t *testing.T) {
	_, err := getArchiveName(opsys.Linux, arch.PPC64)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestGetArchiveName_UnsupportedOS(t *testing.T) {
	_, err := getArchiveName(opsys.Other, arch.PPC64)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func equal(a models.NodeVersion, b models.NodeVersion) bool {
	switch {
	case a.Version != b.Version:
		fallthrough
	case !slices.Equal(a.Files, b.Files):
		fallthrough
	case a.LTS != b.LTS:
		return false
	default:
		return true
	}
}

func slicesEqual(a []models.NodeVersion, b []models.NodeVersion) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if !equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

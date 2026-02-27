package node

import (
	"testing"

	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
)

func TestConvertKeywordToVersion_Invalid(t *testing.T) {
	expected := "invalid"
	actual := convertKeywordToVersion("invalid", opsys.Linux, arch.ARM64, nil, getAllNodeVersionsMock)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertKeywordToVersion_Latest(t *testing.T) {
	expected := "v2.0.0"
	actual := convertKeywordToVersion("latest", opsys.Linux, arch.ARM64, nil, getAllNodeVersionsMock)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertKeywordToVersion_LTS(t *testing.T) {
	expected := "v1.3.5"
	actual := convertKeywordToVersion("lts", opsys.Linux, arch.ARM64, nil, getAllNodeVersionsMock)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertPrefixToVersionDown_Invalid(t *testing.T) {
	_, err := convertPrefixToVersionDown("1.4", opsys.Linux, arch.ARM64, nil, getAllNodeVersionsMock)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestConvertPrefixToVersionDown_Valid(t *testing.T) {
	expected := "v1.3.5"
	actual, err := convertPrefixToVersionDown("1.3", opsys.Linux, arch.ARM64, nil, getAllNodeVersionsMock)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

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

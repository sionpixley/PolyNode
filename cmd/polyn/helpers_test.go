package main

import (
	"os"
	"slices"
	"testing"
	"time"

	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
)

func TestConvertToArchitecture_ARM64(t *testing.T) {
	expected := arch.ARM64
	actual := convertToArchitecture("arm64")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToArchitecture_Other(t *testing.T) {
	expected := arch.Other
	actual := convertToArchitecture("ppc")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToArchitecture_PPC64(t *testing.T) {
	expected := arch.PPC64
	actual := convertToArchitecture("ppc64")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToArchitecture_PPC64LE(t *testing.T) {
	expected := arch.PPC64LE
	actual := convertToArchitecture("ppc64le")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToArchitecture_S390X(t *testing.T) {
	expected := arch.S390X
	actual := convertToArchitecture("s390x")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToArchitecture_X64(t *testing.T) {
	expected := arch.X64
	actual := convertToArchitecture("amd64")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToOperatingSystem_AIX(t *testing.T) {
	expected := opsys.AIX
	actual := convertToOperatingSystem("aix")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToOperatingSystem_Linux(t *testing.T) {
	expected := opsys.Linux
	actual := convertToOperatingSystem("linux")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToOperatingSystem_Mac(t *testing.T) {
	expected := opsys.Mac
	actual := convertToOperatingSystem("darwin")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToOperatingSystem_Other(t *testing.T) {
	expected := opsys.Other
	actual := convertToOperatingSystem("zos")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestConvertToOperatingSystem_Windows(t *testing.T) {
	expected := opsys.Windows
	actual := convertToOperatingSystem("windows")
	if actual != expected {
		t.Errorf("expected: %v actual: %v", expected, actual)
	}
}

func TestDownloadPolyNodeFile(t *testing.T) {
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	err := downloadPolyNodeFile("test", nil, httpWrapper, ioWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case httpWrapper.TimesDoCalled < 1:
		t.Error("expected httpWrapper.Do to have been called\n")
	case ioWrapper.TimesCopyCalled < 1:
		t.Error("expected ioWrapper.Copy to have been called\n")
	case osWrapper.TimesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	case osWrapper.TimesCreateCalled < 1:
		t.Error("expected osWrapper.Create to have been called\n")
	}
}

func TestGetLastUpdate_File(t *testing.T) {
	osWrapper := new(models.OSMockExist)
	expected, err := time.Parse(iso8601, "2025-02-26T12:23:11.723Z")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	actual := getLastUpdate(osWrapper)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetLastUpdate_NoFile(t *testing.T) {
	osWrapper := new(models.OSMockNotExist)
	expected := time.Now().UTC().AddDate(0, 0, -30).Format(iso8601)
	actual := getLastUpdate(osWrapper).Format(iso8601)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestParseCLIArgs_Args(t *testing.T) {
	osWrapper := new(models.OSMockExist)
	os.Args = []string{"polyn", "install", "lts"}
	expected := []string{"install", "lts"}
	actual := parseCLIArgs(osWrapper)
	if len(actual) != 2 {
		t.Errorf("expected: %d actual: %d\n", 2, len(actual))
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

//func TestParseCLIArgs_NoArgs(t *testing.T) {
//	osWrapper := new(models.OSMockExist)
//	os.Args = []string{"polyn"}
//	_ = parseCLIArgs(osWrapper)
//
//	switch {
//	case osWrapper.TimesStdoutCalled < 1:
//		t.Error("expected osWrapper.Stdout to have been called\n")
//	case osWrapper.TimesExitCalled < 1:
//		t.Error("expected osWrapper.Exit to have been called\n")
//	}
//}

//func TestParseCLIArgs_Version(t *testing.T) {
//	osWrapper := new(models.OSMockExist)
//	os.Args = []string{"polyn", "-v"}
//	_ = parseCLIArgs(osWrapper)
//
//	if osWrapper.TimesExitCalled < 1 {
//		t.Error("expected osWrapper.Exit to have been called\n")
//	}
//}

func TestSupportedArchitecture_ARM64(t *testing.T) {
	supported := supportedArchitecture(arch.ARM64)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedArchitecture_PPC64(t *testing.T) {
	supported := supportedArchitecture(arch.PPC64)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedArchitecture_PPC64LE(t *testing.T) {
	supported := supportedArchitecture(arch.PPC64LE)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedArchitecture_S390X(t *testing.T) {
	supported := supportedArchitecture(arch.S390X)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedArchitecture_Unsupported(t *testing.T) {
	supported := supportedArchitecture(arch.Other)
	if supported {
		t.Errorf("expected: %v actual: %v\n", false, supported)
	}
}

func TestSupportedArchitecture_X64(t *testing.T) {
	supported := supportedArchitecture(arch.X64)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedOS_AIX(t *testing.T) {
	supported := supportedOS(opsys.AIX)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedOS_Linux(t *testing.T) {
	supported := supportedOS(opsys.Linux)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedOS_Mac(t *testing.T) {
	supported := supportedOS(opsys.Mac)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestSupportedOS_Unsupported(t *testing.T) {
	supported := supportedOS(opsys.Other)
	if supported {
		t.Errorf("expected: %v actual: %v\n", false, supported)
	}
}

func TestSupportedOS_Windows(t *testing.T) {
	supported := supportedOS(opsys.Windows)
	if !supported {
		t.Errorf("expected: %v actual: %v\n", true, supported)
	}
}

func TestRunUpdateScript_AIX(t *testing.T) {
	execWrapper := new(models.ExecMock)
	osWrapper := new(models.OSMockExist)
	err := runUpdateScript(opsys.AIX, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.TimesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.TimesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Linux(t *testing.T) {
	execWrapper := new(models.ExecMock)
	osWrapper := new(models.OSMockExist)
	err := runUpdateScript(opsys.Linux, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.TimesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.TimesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Mac(t *testing.T) {
	execWrapper := new(models.ExecMock)
	osWrapper := new(models.OSMockExist)
	err := runUpdateScript(opsys.Mac, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.TimesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.TimesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Windows(t *testing.T) {
	execWrapper := new(models.ExecMock)
	osWrapper := new(models.OSMockExist)
	err := runUpdateScript(opsys.Windows, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case osWrapper.TimesWriteFileCalled < 1:
		t.Error("expected osWrapper.WriteFile to have been called\n")
	case execWrapper.TimesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	}
}

func TestUpdatePolyNode_AIX(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(opsys.AIX, arch.PPC64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestUpdatePolyNode_Linux(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(opsys.Linux, arch.X64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestUpdatePolyNode_Mac(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(opsys.Mac, arch.ARM64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestUpdatePolyNode_UnsupportedArch(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(opsys.Linux, arch.PPC64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestUpdatePolyNode_UnsupportedOS(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(12, arch.X64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err == nil {
		t.Error("expected: error actual: nil\n")
	}
}

func TestUpdatePolyNode_Windows(t *testing.T) {
	execWrapper := new(models.ExecMock)
	gzipWrapper := new(models.GzipMock)
	httpWrapper := new(models.HTTPMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	zipWrapper := new(models.ZipMock)

	err := updatePolyNode(opsys.Windows, arch.X64, nil, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper, zipWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

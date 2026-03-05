package main

import (
	"testing"
	"time"

	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
)

const dateFormat = "2006-01-02"

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
	httpWrapper := new(httpMock)
	ioWrapper := new(ioMock)
	osWrapper := new(osMockExist)
	err := downloadPolyNodeFile("test", httpWrapper, ioWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case httpWrapper.timesDoCalled < 1:
		t.Error("expected httpWrapper.Do to have been called\n")
	case ioWrapper.timesCopyCalled < 1:
		t.Error("expected ioWrapper.Copy to have been called\n")
	case osWrapper.timesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	case osWrapper.timesCreateCalled < 1:
		t.Error("expected osWrapper.Create to have been called\n")
	}
}

func TestGetLastUpdate_NoFile(t *testing.T) {
	osWrapper := new(osMockNotExist)
	expected := time.Now().UTC().AddDate(0, 0, -30).Format(dateFormat)
	actual := getLastUpdate(osWrapper).Format(dateFormat)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestGetLastUpdate_File(t *testing.T) {
	osWrapper := new(osMockExist)
	expected, err := time.Parse(isoDateTimeFormat, "2025-02-26T12:23:11.723Z")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	actual := getLastUpdate(osWrapper)
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestRunUpdateScript_AIX(t *testing.T) {
	execWrapper := new(execMock)
	osWrapper := new(osMockExist)
	err := runUpdateScript(opsys.AIX, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.timesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.timesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Linux(t *testing.T) {
	execWrapper := new(execMock)
	osWrapper := new(osMockExist)
	err := runUpdateScript(opsys.Linux, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.timesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.timesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Mac(t *testing.T) {
	execWrapper := new(execMock)
	osWrapper := new(osMockExist)
	err := runUpdateScript(opsys.Mac, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case execWrapper.timesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	case osWrapper.timesRemoveAllCalled < 1:
		t.Error("expected osWrapper.RemoveAll to have been called\n")
	}
}

func TestRunUpdateScript_Windows(t *testing.T) {
	execWrapper := new(execMock)
	osWrapper := new(osMockExist)
	err := runUpdateScript(opsys.Windows, execWrapper, osWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case osWrapper.timesWriteFileCalled < 1:
		t.Error("expected osWrapper.WriteFile to have been called\n")
	case execWrapper.timesRunCalled < 1:
		t.Error("expected execWrapper.Run to have been called\n")
	}
}

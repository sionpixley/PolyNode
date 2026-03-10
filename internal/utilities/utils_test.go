package utilities

import (
	"errors"
	"testing"

	"github.com/sionpixley/PolyNode/internal/constants/command"
	"github.com/sionpixley/PolyNode/internal/models"
)

func TestConvertToCommand_Add(t *testing.T) {
	expected := command.Add
	actual := ConvertToCommand("add")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_ConfigGet(t *testing.T) {
	expected := command.ConfigGet
	actual := ConvertToCommand("config-get")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_ConfigSet(t *testing.T) {
	expected := command.ConfigSet
	actual := ConvertToCommand("config-set")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Current(t *testing.T) {
	expected := command.Current
	actual := ConvertToCommand("current")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Default(t *testing.T) {
	expected := command.Default
	actual := ConvertToCommand("default")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Install(t *testing.T) {
	expected := command.Install
	actual := ConvertToCommand("install")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_List(t *testing.T) {
	expected := command.List
	actual := ConvertToCommand("list")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Ls(t *testing.T) {
	expected := command.List
	actual := ConvertToCommand("ls")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Other(t *testing.T) {
	expected := command.Other
	actual := ConvertToCommand("asdfhapwueifj")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Remove(t *testing.T) {
	expected := command.Remove
	actual := ConvertToCommand("remove")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Rm(t *testing.T) {
	expected := command.Remove
	actual := ConvertToCommand("rm")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Search(t *testing.T) {
	expected := command.Search
	actual := ConvertToCommand("search")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToCommand_Use(t *testing.T) {
	expected := command.Use
	actual := ConvertToCommand("use")
	if actual != expected {
		t.Errorf("expected: %v actual: %v\n", expected, actual)
	}
}

func TestConvertToSemanticVersion_WithoutV(t *testing.T) {
	expected := "v2.1.56"
	actual := ConvertToSemanticVersion("2.1.56")
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestConvertToSemanticVersion_WithV(t *testing.T) {
	expected := "v2.1.56"
	actual := ConvertToSemanticVersion("v2.1.56")
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestExtractGzip(t *testing.T) {
	gzipWrapper := new(models.GzipMock)
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	tarWrapper := new(models.TarMock)
	err := ExtractGzip("test.tar.gz", "node", gzipWrapper, ioWrapper, osWrapper, tarWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case osWrapper.TimesOpenCalled < 1:
		t.Error("expected osWrapper.Open to have been called\n")
	case gzipWrapper.TimesNewReaderCalled < 1:
		t.Error("expected gzipWrapper.NewReader to have been called\n")
	case tarWrapper.TimesNewReaderCalled < 1:
		t.Error("expected tarWrapper.NewReader to have been called\n")
	case tarWrapper.TimesNextCalled < 1:
		t.Error("expected tarWrapper.Next to have been called\n")
	case gzipWrapper.TimesCloseCalled < 1:
		t.Error("expected gzipWrapper.Close to have been called\n")
	}
}

func TestExtractZip(t *testing.T) {
	ioWrapper := new(models.IOMock)
	osWrapper := new(models.OSMockExist)
	zipWrapper := new(models.ZipMock)
	err := ExtractZip("test.zip", "node", ioWrapper, osWrapper, zipWrapper)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	switch {
	case zipWrapper.TimesOpenReaderCalled < 1:
		t.Error("expected zipWrapper.OpenReader to have been called\n")
	case zipWrapper.TimesFileCalled < 1:
		t.Error("expected zipWrapper.File to have been called\n")
	case zipWrapper.TimesCloseCalled < 1:
		t.Error("expected zipWrapper.Close to have been called\n")
	}
}

func TestKnownCommand_Known(t *testing.T) {
	known := KnownCommand("use")
	if !known {
		t.Errorf("expected: %v actual: %v\n", true, known)
	}
}

func TestKnownCommand_Unknown(t *testing.T) {
	known := KnownCommand("unknown")
	if known {
		t.Errorf("expected: %v actual: %v\n", false, known)
	}
}

func TestLogUserError(t *testing.T) {
	osWrapper := new(models.OSMockExist)
	LogUserError(errors.New("test"), osWrapper)

	switch {
	case osWrapper.TimesStderrCalled < 2:
		t.Error("expected osWrapper.Stderr to have been called\n")
	case osWrapper.TimesExitCalled < 1:
		t.Error("expected osWrapper.Exit to have been called\n")
	}
}

func TestStripTopDir_OnePart(t *testing.T) {
	expected := "example"
	actual := stripTopDir("example")
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestStripTopDir_TwoParts(t *testing.T) {
	expected := "example"
	actual := stripTopDir("idk/example")
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestStripTopDir_ThreeParts(t *testing.T) {
	expected := "example/hello"
	actual := stripTopDir("idk/example/hello")
	if actual != expected {
		t.Errorf("expected: %s actual: %s\n", expected, actual)
	}
}

func TestValidVersionFormat_InvalidCharacter(t *testing.T) {
	valid := ValidVersionFormat("v3.1b")
	if valid {
		t.Errorf("expected: %v actual: %v\n", false, valid)
	}
}

func TestValidVersionFormat_InvalidNotEnoughParts(t *testing.T) {
	valid := ValidVersionFormat("v3.1")
	if valid {
		t.Errorf("expected: %v actual: %v\n", false, valid)
	}
}

func TestValidVersionFormat_ValidWithoutV(t *testing.T) {
	valid := ValidVersionFormat("3.1.0")
	if !valid {
		t.Errorf("expected: %v actual: %v\n", true, valid)
	}
}

func TestValidVersionFormat_ValidWithV(t *testing.T) {
	valid := ValidVersionFormat("v3.1.0")
	if !valid {
		t.Errorf("expected: %v actual: %v\n", true, valid)
	}
}

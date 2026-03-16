package main

import (
	"testing"

	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
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

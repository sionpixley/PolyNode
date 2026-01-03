package arch

import "github.com/sionpixley/PolyNode/internal/models"

const (
	Other models.Architecture = iota
	ARM64
	PPC64
	PPC64LE
	S390X
	X64
)

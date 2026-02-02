package command

import "github.com/sionpixley/PolyNode/internal/models"

// We don't include the version command in this. The version command is handled in main().
// We don't include the update command in this either. It also gets handled in main().
const (
	Other models.Command = iota
	Add
	Current
	Default
	Install
	List
	Remove
	Search
	Use
)

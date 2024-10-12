package constants

import "github.com/sionpixley/PolyNode/internal/models"

const (
	HELP string = `
Description:

  PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

Usage:

  polyn command

Commands:

  add <version>     : Downloads a specific version of Node.js, but does not switch it to your current version.
  add lts           : Downloads the latest LTS version of Node.js.
  add latest        : Downloads the latest version of Node.js.
  current           : Prints the current version of Node.js.
  install <version> : Downloads a specific version of Node.js and sets it as your current version.
  install lts       : Downloads the latest LTS version of Node.js and sets it as your current version.
  install latest    : Downloads the latest version of Node.js and sets it as your current version.
  list              : Prints the list of downloaded Node.js versions.
  ls                : Alias for list command.
  remove <version>  : Deletes a version of Node.js.
  rm <version>      : Alias for remove command.
  search            : Prints out a list of the most recent Node.js versions available for download.
  search <prefix>   : Prints out a list of Node.js versions that have this prefix.
  use <version>     : Switches Node.js to a different version.
  upgrade           : Upgrades PolyNode to the latest release.
  version           : Prints the current version of PolyNode.`

	INVALID_VERSION_FORMAT_ERROR string = "invalid version format"
	UNSUPPORTED_OS_ERROR         string = "unsupported operating system"

	// PolyNode's version.
	VERSION string = "v0.11.0"
)

// NA is for Not Applicable.
const (
	NA_OS models.OperatingSystem = iota
	AIX
	LINUX
	MAC
	WINDOWS
)

// NA is for Not Applicable.
const (
	NA_ARCH models.Architecture = iota
	ARM64
	PPC64
	X64
)

// NA is for Not Applicable.
// We don't include the version command in this. The version command is handled in main().
const (
	NA_COMM models.Command = iota
	ADD
	CURRENT
	INSTALL
	LIST
	REMOVE
	SEARCH
	UPGRADE
	USE
)

package internal

const (
	HELP string = `
Description:

  PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

Usage:

  polyn command

Commands:

  add <version>     : Downloads a specific version of Node.js, but does not switch it to your current version.
  current           : Prints the current version of Node.js.
  install <version> : Equivalent to the add command followed by the use command.
  list              : Prints the list of downloaded Node.js versions.
  ls                : Alias for list command.
  remove <version>  : Deletes a version of Node.js.
  rm <version>      : Alias for remove command.
  search            : Prints out a list of the most recent Node.js versions available for download.
  search <prefix>   : Prints out a list of Node.js versions that have this prefix.
  use <version>     : Switches Node.js to a different version.
  version           : Prints the current version of PolyNode.`

	UNSUPPORTED_OS_ERROR string = "unsupported operating system"

	// PolyNode's version.
	VERSION string = "v0.6.2"

	_DEFAULT_NODE_MIRROR          string = "https://nodejs.org/dist"
	_INVALID_VERSION_FORMAT_ERROR string = "invalid version format"
)

// NA is for Not Applicable.
const (
	NA_OS OperatingSystem = iota
	LINUX
	MAC
	WINDOWS
)

// NA is for Not Applicable.
const (
	NA_ARCH Architecture = iota
	ARM64
	X64
)

// NA is for Not Applicable.
// We don't include the version command in this. The version command is handled in main().
const (
	_NA_COMM command = iota
	_ADD
	_CURRENT
	_INSTALL
	_LIST
	_REMOVE
	_SEARCH
	_USE
)

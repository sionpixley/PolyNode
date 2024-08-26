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
  use <version>     : Switches Node.js to a different version.
  version           : Prints the current version of PolyNode.`

	// PolyNode's version.
	VERSION string = "v0.2.1"

	// Unsupported operating system error message.
	c_UNSUPPORTED_OS string = "unsupported operating system"
)

// NA is for Not Applicable.
const (
	c_NA_ARCH Architecture = iota
	c_ARM64
	c_X64
)

// NA is for Not Applicable.
// We don't include the version command in this. The version command is handled in main().
const (
	c_NA_COMM command = iota
	c_ADD
	c_CURRENT
	c_INSTALL
	c_LIST
	c_REMOVE
	c_SEARCH
	c_USE
)

// NA is for Not Applicable.
const (
	c_NA_OS OperatingSystem = iota
	c_LINUX
	c_MAC
)

package internal

const (
	// List of commands for printing to the user.
	c_COMMANDS string = `Commands:

  add <version>     : Downloads a specific version of Node.js, but does not switch it to your current version.
  current           : Prints the current version of Node.js.
  install <version> : Downloads a version of Node.js and switches it to your current version.
  list              : Prints the list of downloaded Node.js versions.
  ls                : Alias for list command.
  remove <version>  : Deletes a version of Node.js.
  rm <version>      : Alias for remove command.
  search            : Prints out a list of the most recent Node.js versions available for download.
  use <version>     : Switches Node.js to a different version.
  version           : Prints the current version of PolyNode.`

	// PolyNode description for printing to the user.
	c_DESCRIPTION string = `Description:

  PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.`

	// Unsupported operating system error message.
	c_UNSUPPORTED_OS string = "unsupported operating system"

	// Syntax for using the polyn CLI.
	c_USAGE string = `Usage:

  polyn command`

	// PolyNode's version.
	VERSION string = "v0.2.0"
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

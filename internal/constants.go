package internal

const (
	// List of commands for printing to the user.
	c_COMMANDS string = `Commands:
	
      add : Downloads a version of the env, but does not switch it to your current version.
  current : Prints the current version of the env.
  install : Downloads a version of the env and switches it to your current version.
     list : Prints the list of downloaded versions for the env.
       ls : Alias for list command.
   remove : Deletes a version of the env.
       rm : Alias for remove command.
      use : Switches the env to a different version.
  version : Prints the current version of PolyNode.`

	// PolyNode description for printing to the user.
	c_DESCRIPTION string = `Description:

  PolyNode is a CLI tool that helps install and switch between multiple versions of Bun, Deno, and/or Node.js on the same device.`

	// List of runtimes for printing to the user.
	c_RUNTIMES string = `Runtimes:

  bun
  deno
  node : Default if runtime is not provided.`

	// Unsupported operating system error message.
	c_UNSUPPORTED_OS string = "unsupported operating system"

	// Syntax for using the polyn CLI.
	c_USAGE string = `Usage:

  polyn [runtime] command <version>`

	// PolyNode's version.
	VERSION string = "v0.1.0"
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
	c_USE
)

// NA is for Not Applicable.
const (
	c_NA_OS OperatingSystem = iota
	c_LINUX
	c_MAC
	c_WIN
)

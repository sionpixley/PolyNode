package constants

import "github.com/sionpixley/polyn/internal/models"

const (
	// List of commands for printing to the user.
	COMMANDS string = `Commands:
	
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
	DESCRIPTION string = `
Description:

  PolyNode is a CLI tool that helps install and switch between multiple versions of Bun, Deno, and/or Node.js on the same device. It is compatible with Linux, macOS, and Windows.`

	// List of envs for printing to the user.
	ENVS string = `Envs:

  bun
  deno
  node : Default if env is not provided.`

	// Syntax for using the polyn CLI.
	USAGE string = "Usage:\n\n  polyn [env] command <version>"

	// PolyNode's version.
	VERSION string = "v0.1.0"
)

// NA is for Not Applicable.
// We don't include the version command in this. The version command is handled in main().
const (
	NA models.Command = iota
	ADD
	CURRENT
	INSTALL
	LIST
	REMOVE
	USE
)

func ConvertToCommand(commandStr string) models.Command {
	switch commandStr {
	case "add":
		return ADD
	case "current":
		return CURRENT
	case "install":
		return INSTALL
	case "ls":
		fallthrough
	case "list":
		return LIST
	case "rm":
		fallthrough
	case "remove":
		return REMOVE
	case "use":
		return USE
	default:
		return NA
	}
}

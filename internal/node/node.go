package node

import (
	"fmt"
	"os/exec"

	"github.com/sionpixley/polyn/internal"
	"github.com/sionpixley/polyn/internal/constants"
)

// Main function for Node.js actions.
// The args parameter should not include the optional env.
func Handle(args []string) {
	if args == nil || len(args) == 0 {
		internal.PrintHelp()
		return
	}

	var err error
	command := constants.ConvertToCommand(args[0])
	switch command {
	case constants.ADD:
		if len(args) > 1 {
			err = add(args[1])
		} else {
			internal.PrintHelp()
		}
	case constants.CURRENT:
		printCurrent()
	case constants.INSTALL:
		if len(args) > 1 {
			err = install(args[1])
		} else {
			internal.PrintHelp()
		}
	case constants.LIST:
		listDownloaded()
	case constants.REMOVE:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			internal.PrintHelp()
		}
	case constants.USE:
		if len(args) > 1 {
			err = use(args[1])
		} else {
			internal.PrintHelp()
		}
	default:
		internal.PrintHelp()
	}

	if err != nil {
		internal.PrintError(err)
	}
}

func add(version string) error {
	return nil
}

func install(version string) error {
	err := add(version)
	if err != nil {
		return err
	}

	err = use(version)
	return err
}

func listDownloaded() {

}

func printCurrent() {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Print(string(output))
	}
}

func remove(version string) error {
	return nil
}

func use(version string) error {
	return nil
}

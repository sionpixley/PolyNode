package node

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	command := internal.ConvertToCommand(args[0])
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
	fileName := "node-" + version + "-win-x64.zip"
	fmt.Println("Downloading " + fileName + "...")

	url := "https://nodejs.org/dist/" + version + "/" + fileName

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = os.Mkdir("node", os.ModePerm)
	if err != nil {
		return err
	}

	filePath := "./node/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	output, err := exec.Command("./emb/7za.exe", "x", ".\\node\\node-v20.13.1-win-x64.zip", "-o.\\node\\"+version).Output()
	if err != nil {
		return err
	}
	fmt.Print(output)

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

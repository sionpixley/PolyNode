package main

import (
	"os"

	"github.com/sionpixley/PolyNode/internal/models"
)

func main() {
	operatingSystem := checkOS()
	arch := checkArchitecture()
	config := models.NewPolyNodeConfig()
	args := parseCLIArgs()

	execute(
		args,
		operatingSystem,
		arch,
		config,
		os.IsNotExist,
		os.ReadFile,
		os.Stat,
	)
}

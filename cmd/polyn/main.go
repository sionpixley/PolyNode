package main

import (
	"github.com/sionpixley/PolyNode/internal/models"
)

func main() {
	operatingSystem := checkOS()
	arch := checkArchitecture()
	config := models.LoadPolyNodeConfig()
	args := parseCLIArgs()

	execute(args, operatingSystem, arch, config)
}

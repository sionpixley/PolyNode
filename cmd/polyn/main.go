package main

import (
	"github.com/sionpixley/PolyNode/internal/models"
)

func main() {
	var osWrapper models.OSWrap

	operatingSystem := checkOS()
	arch := checkArchitecture()
	config := models.NewPolyNodeConfig()
	args := parseCLIArgs()

	execute(args, operatingSystem, arch, config, osWrapper)
}

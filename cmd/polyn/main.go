package main

import (
	"github.com/sionpixley/PolyNode/internal/models"
)

func main() {
	var httpWrapper models.HTTPWrap
	var ioWrapper models.IOWrap
	var osWrapper models.OSWrap

	operatingSystem := checkOS()
	arch := checkArchitecture()
	config := models.NewPolyNodeConfig()
	args := parseCLIArgs(osWrapper)

	execute(args, operatingSystem, arch, config, httpWrapper, ioWrapper, osWrapper)
}

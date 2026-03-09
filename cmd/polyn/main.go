package main

import (
	"github.com/sionpixley/PolyNode/internal/models"
)

func main() {
	var execWrapper models.ExecWrap
	var gzipWrapper models.GzipWrap
	var httpWrapper models.HTTPWrap
	var ioWrapper models.IOWrap
	var osWrapper models.OSWrap
	var tarWrapper models.TarWrap

	operatingSystem := checkOS()
	arch := checkArchitecture()
	config := models.NewPolyNodeConfig(osWrapper)
	args := parseCLIArgs(osWrapper)

	execute(args, operatingSystem, arch, config, execWrapper, gzipWrapper, httpWrapper, ioWrapper, osWrapper, tarWrapper)
}

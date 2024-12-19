package constants

import "github.com/sionpixley/PolyNode/internal/models"

const (
	HELP string = `
Description:

    PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

Usage:

    polyn command

Commands:

    add <version or keyword>
        Downloads a specific version of Node.js, but does not switch it to your current version.
    current
        Prints your current version of Node.js.
    install <version or keyword>
        Downloads a specific version of Node.js and sets it as your default version.
    list
        Prints the list of downloaded Node.js versions.
    ls
        Alias for 'list' command.
    remove <version>
        Deletes a version of Node.js.
    rm <version>
        Alias for 'remove' command.
    search [prefix]
        Prints out a list of Node.js versions that have this prefix.
        If the prefix is omitted, prints out a list of the most recent Node.js versions available for download.
    temp <version>
        Prints out the command needed to temporarily set your Node.js to a specific version.
        If on AIX, Linux, or macOS, please use 'eval $(polyn temp <version>)' instead.
    use <version>
        Sets your default Node.js to a different version.
    upgrade
        Upgrades PolyNode to the latest release.
    version
        Prints the current version of PolyNode.

Keywords:

    latest
        Represents the latest release of Node.js.
    lts
        Represents the most recent LTS release of Node.js.`

	INVALID_VERSION_FORMAT_ERROR string = "polyn: invalid version format"
	UNSUPPORTED_ARCH_ERROR       string = "polyn: unsupported CPU architecture"
	UNSUPPORTED_OS_ERROR         string = "polyn: unsupported operating system"

	// PolyNode's version.
	VERSION string = "v2.0.5"
)

// NA is for Not Applicable.
const (
	NA_OS models.OperatingSystem = iota
	AIX
	LINUX
	MAC
	WINDOWS
)

// NA is for Not Applicable.
const (
	NA_ARCH models.Architecture = iota
	ARM64
	PPC64
	PPC64LE
	S390X
	X64
)

// NA is for Not Applicable.
// We don't include the version command in this. The version command is handled in main().
// We don't include the upgrade command in this either. It also gets handled in main().
const (
	NA_COMM models.Command = iota
	ADD
	CURRENT
	INSTALL
	LIST
	REMOVE
	SEARCH
	TEMP
	USE
)

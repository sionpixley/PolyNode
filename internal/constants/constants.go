package constants

import "github.com/sionpixley/PolyNode/internal/models"

const (
	Help string = `Description:

    PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

Usage:

    polyn command

Commands:

    add <version or keyword or prefix>
        Downloads a specific version of Node.js, but does not switch it to your current version.
    current
        Prints your current version of Node.js.
    install <version or keyword or prefix>
        Downloads a specific version of Node.js and sets it as your default version.
    list
        Prints the list of downloaded Node.js versions.
    ls
        Alias for 'list' command.
    remove <version or prefix>
        Deletes a version of Node.js.
    rm <version or prefix>
        Alias for 'remove' command.
    search [prefix]
        Prints out a list of Node.js versions that have this prefix.
        If the prefix is omitted, prints out a list of the most recent Node.js versions available for download.
    temp <version or prefix>
        Prints out the command needed to temporarily set your Node.js to a specific version.
        If on AIX, Linux, or macOS, please use 'eval $(polyn temp <version>)' instead.
    use <version or prefix>
        Sets your default Node.js to a different version.
    update
        Updates PolyNode to the latest release.
    version
        Prints the current version of PolyNode.

Keywords:

    latest
        Represents the latest release of Node.js.
    lts
        Represents the most recent LTS release of Node.js.`

	NoDownloadedNodejsMessage string = "There are no Node.js versions downloaded.\nTo download a Node.js version, use the 'add' or 'install' command."

	UnsupportedArchError string = "polyn error: unsupported CPU architecture"
	UnsupportedOSError   string = "polyn error: unsupported operating system"

	// PolyNode's version.
	Version string = "v3.0.5"
)

const (
	OtherOS models.OperatingSystem = iota
	Aix
	Linux
	Mac
	Windows
)

const (
	OtherArch models.Architecture = iota
	Arm64
	Ppc64
	Ppc64Le
	S390x
	X64
)

// We don't include the version command in this. The version command is handled in main().
// We don't include the update command in this either. It also gets handled in main().
const (
	OtherComm models.Command = iota
	Add
	Current
	Install
	List
	Remove
	Search
	Temp
	Use
)

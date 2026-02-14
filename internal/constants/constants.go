package constants

const (
	Help string = `Description:

    PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

Usage:

    polyn [option...] <command>

Commands:

    add <version | keyword | prefix>
        Downloads a specific version of Node.js, but does not switch it to your current version.
        Prefix will match the newest version with that prefix.
    config-get [config_field]
        Prints the current value for the config field.
        If the config field is omitted, the full configuration file is printed.
    config-set <config_field> <value>
        Sets the value for the config field.
    current
        Prints your current version of Node.js.
    default <version | prefix>
        Sets your default Node.js to a different version.
        Prefix will match the newest version with that prefix.
    install <version | keyword | prefix>
        Downloads a specific version of Node.js and sets it as your default version.
        Prefix will match the newest version with that prefix.
    list
        Prints the list of downloaded Node.js versions.
    ls
        Alias for 'list' command.
    remove <version | prefix>
        Deletes a version of Node.js.
        Prefix will match the oldest version with that prefix.
    rm <version | prefix>
        Alias for 'remove' command.
        Prefix will match the oldest version with that prefix.
    search [prefix]
        Prints out a list of Node.js versions that have this prefix.
        If the prefix is omitted, prints out a list of the most recent Node.js versions available for download.
    update
        Updates PolyNode to the latest release.
    use <version | prefix>
        Prints out the command needed to temporarily set your Node.js to a specific version.
        If on AIX, Linux, or macOS, please use 'eval $(polyn use <version or prefix>)' instead.
        Prefix will match the newest version with that prefix.

Options:

    -h, --help
        Prints help and usage information.
    -v, --version
        Prints the current version of PolyNode.

Keywords:

    latest
        Represents the latest release of Node.js.
    lts
        Represents the most recent LTS release of Node.js.

Config fields:

    autoUpdate
        Bool that configures if PolyNode's auto updater should run.
        Default value is 'true'.
    nodeMirror
        String that configures the URL to download Node.js.
        Default value is 'https://nodejs.org/dist'.`

	InvalidConfigFieldError            string = "invalid config field: '%s'"
	MissingVersionKeywordOrPrefixError string = "missing argument: the '%s' command is missing a version, keyword, or prefix"
	MissingVersionOrPrefixError        string = "missing argument: the '%s' command is missing a version or prefix"

	NoDownloadedNodejsMessage string = "There are no Node.js versions downloaded.\nTo download a Node.js version, use the 'add' or 'install' command."

	UnknownCommandError  string = "unknown command: '%s' is not a known command"
	UnsupportedArchError string = "polyn: unsupported CPU architecture"
	UnsupportedOSError   string = "polyn: unsupported operating system"

	// Version constant is PolyNode's version.
	Version string = "v4.0.2"
)

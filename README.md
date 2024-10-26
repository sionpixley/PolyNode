# PolyNode

PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device. It does not require sudo/admin privileges, and is installed on a per-user basis.

PolyNode has a GUI that you can use, but it must be installed first. Release assets prefixed with `PolyNode-GUI` will install the GUI along with the base CLI command, `polyn`. Read about [launching the GUI](#launching-the-gui) below.

## Table of contents

1. [Supported operating systems and CPU architectures](#supported-operating-systems-and-cpu-architectures)
2. [How to install](#how-to-install-polynode)
    1. [AIX](#for-aix)
    2. [Linux](#for-linux)
    3. [macOS](#for-macos)
    4. [Windows](#for-windows)
3. [How to use](#how-to-use)
    1. [Launching the GUI](#launching-the-gui)
    2. [Upgrading PolyNode](#upgrading-polynode-to-the-latest-release)
    3. [Searching for available Node.js versions](#searching-for-available-nodejs-versions)
    4. [Searching for a specific Node.js version](#searching-for-a-specific-nodejs-version)
    5. [Downloading and switching to a new version of Node.js](#downloading-and-switching-to-a-new-version-of-nodejs)
    6. [Downloading a new version of Node.js](#downloading-a-new-version-of-nodejs)
    7. [Switching to a different downloaded version of Node.js](#switching-to-a-different-downloaded-version-of-nodejs)
    8. [Printing your current version of Node.js](#printing-your-current-version-of-nodejs)
    9. [Printing all downloaded versions of Node.js](#printing-all-downloaded-versions-of-nodejs)
    10. [Deleting a downloaded version of Node.js](#deleting-a-downloaded-version-of-nodejs)
    11. [Printing your current version of PolyNode](#printing-your-current-version-of-polynode)
4. [How to configure](#how-to-configure-polynode)
    1. [Configuration fields](#configuration-fields)
5. [How to uninstall](#how-to-uninstall-polynode)
    1. [AIX, Linux, and macOS](#aix-linux-and-macos)
    2. [Windows](#windows)
6. [Building from source](#building-from-source)
    1. [Required technologies](#required-technologies)
    2. [AIX](#building-on-aix)
    3. [Linux](#building-on-linux)
    4. [macOS](#building-on-macos)
    5. [Windows](#building-on-windows)
7. [Future development](#future-development)
8. [Acknowledgements](#acknowledgements)

## Supported operating systems and CPU architectures

- AIX (Power ISA 64-bit)
- Linux (x64 and ARM64)
- macOS (x64 and ARM64)
- Windows 10 and newer (x64 and ARM64)

## How to install PolyNode

PolyNode does not require sudo/admin privileges to install.

Please uninstall all Node.js downloads that weren't installed by PolyNode before running the setup binary.

### For AIX

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the AIX .tar.gz file.
3. Extract the .tar.gz file and run the setup binary.

### For Linux

PolyNode only supports Bash, Zsh, or KornShell by default. During the install process, PolyNode edits either .bashrc, .zshrc, or .kshrc to add two locations to the PATH: PolyNode's home directory `$HOME/.PolyNode` and the symlink for Node.js `$HOME/.PolyNode/nodejs/bin`. You can get PolyNode to work for other shells by adding these directories to your PATH environment variable.

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the Linux .tar.xz file appropriate for your CPU architecture.
3. Extract the .tar.xz file and run the setup binary.

### For macOS

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the Darwin .tar.gz file appropriate for your CPU architecture.
3. Extract the .tar.gz file and run the setup binary.

### For Windows

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the Windows .zip file appropriate for your CPU architecture.
3. Extract the .zip file and run setup.exe.

## How to use

PolyNode does not require sudo/admin privileges to use the `polyn` nor the `PolyNode` command.

### Launching the GUI

If you installed PolyNode's GUI, type this command into your terminal:

`PolyNode`

### Upgrading PolyNode to the latest release

`polyn upgrade`

### Searching for available Node.js versions

`polyn search`

### Searching for a specific Node.js version

Using a prefix will match anything with this prefix. So `polyn search 1` will match with any Node.js version that starts with "1".

If you want to search for a specific major version, add a "." at the end of your prefix. `polyn search 18.` will print all Node.js v18 releases.

`polyn search <prefix>`

### Downloading and switching to a new version of Node.js

The `install` command is equivalent to the `add` command followed by the `use` command.

`polyn install <version>`

You can also use the `lts` keyword to download the latest LTS release without providing a specific version. 

`polyn install lts`

The `latest` keyword will download the latest release of Node.js, regardless if it's an LTS version or not.

`polyn install latest`

### Downloading a new version of Node.js

`polyn add <version>`

`polyn add lts`

`polyn add latest`

### Switching to a different downloaded version of Node.js

`polyn use <version>`

### Printing your current version of Node.js

`polyn current`

### Printing all downloaded versions of Node.js

`polyn ls`

or 

`polyn list`

### Deleting a downloaded version of Node.js

`polyn rm <version>`

or 

`polyn remove <version>`

### Printing your current version of PolyNode

`polyn version`

## How to configure PolyNode

PolyNode's configuration is handled through a JSON file named `polynrc.json` located in PolyNode's home directory (`$HOME/.PolyNode` for Linux/macOS and `%LOCALAPPDATA%\Programs\PolyNode` for Windows). Please see below for the default configuration for `polynrc.json`:

```
{
  "guiPort": 2334,
  "nodeMirror": "https://nodejs.org/dist"
}
```

### Configuration fields

#### guiPort

This field is an `int` that represents the port number the GUI binds to when launched. Default value is `2334`. Only change it if you have another process using that port.

#### nodeMirror

This field is a `string` that represents the URL to download Node.js. Default value is `"https://nodejs.org/dist"`.

## How to uninstall PolyNode

PolyNode does not require sudo/admin privileges to uninstall.

### AIX, Linux, and macOS

1. Run the `$HOME/.PolyNode/uninstall/uninstall` binary.

### Windows

1. Run `%LOCALAPPDATA%\Programs\PolyNode\uninstall\uninstall.exe`.

## Building from source

### Required technologies

- Go 1.23.2
- Node.js 20.18.0 (if building GUI)
- Angular 18.2.8 (if building GUI)
- pnpm 9.12.2 (if building GUI)

### Building on AIX

Run the POSIX shell script `./scripts/aix/bundle`. This script will build PolyNode's source code for Power ISA 64-bit (with and without the GUI), and bundle the artifacts as separate .tar.gz files.

### Building on Linux

Run the POSIX shell script `./scripts/linux/bundle`. This script will build PolyNode's source code for x64 and ARM64 (with and without the GUI), and bundle the artifacts as separate .tar.xz files.

### Building on macOS

macOS has a POSIX shell script (`./scripts/mac/bundle`) that builds and notarizes PolyNode's source code for x64 and ARM64 (with and without the GUI), and bundles the artifacts as separate .tar.gz files. If you don't need to distribute the binaries, then you don't need the notarization step. Just edit the bundle script and set the `sign` variable to `0`.

### Building on Windows

Run the batchfile `.\scripts\win\bundle.cmd`. This batchfile will build PolyNode's source code for x64 and ARM64 (with and without the GUI), and bundle the artifacts as separate .zip files.

## Future development

The original scope of this project was to be able to install and manage multiple versions of Bun, Deno, and Node.js. It currently only supports Node.js, but I would like to support Bun and Deno in the future.

## Acknowledgements

PolyNode draws a lot of inspiration, especially in regards to syntax, from other, more well-known projects, like: [nvm](https://github.com/nvm-sh/nvm), [nvm-windows](https://github.com/coreybutler/nvm-windows), and [nvs](https://github.com/jasongin/nvs).

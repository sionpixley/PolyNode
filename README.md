# PolyNode

PolyNode is a Node.js version manager that does not require sudo/admin privileges and is installed on a per-user basis. Works on AIX, Linux, macOS, and Windows.

## Table of contents

1. [Quick demo](#quick-demo)
2. [Supported operating systems and CPU architectures](#supported-operating-systems-and-cpu-architectures)
3. [How to install](#how-to-install-polynode)
    1. [AIX](#for-aix)
    2. [Linux](#for-linux)
    3. [macOS](#for-macos)
    4. [Windows](#for-windows)
4. [How to use](#how-to-use)
    1. [Updating PolyNode](#updating-polynode-to-the-latest-release)
    2. [Searching for available Node.js versions](#searching-for-available-nodejs-versions)
    3. [Downloading a new version of Node.js](#downloading-a-new-version-of-nodejs)
        1. [Examples](#examples)
    4. [Setting your default Node.js version](#setting-your-default-nodejs-version)
        1. [Examples](#examples-1)
    5. [Temporarily setting your Node.js version](#temporarily-setting-your-nodejs-version)
        1. [AIX, Linux, or macOS](#temporarily-setting-your-nodejs-on-aix-linux-or-macos)
            1. [Examples](#examples-2)
        2. [Windows](#temporarily-setting-your-nodejs-on-windows)
    6. [Downloading and setting your default Node.js to a new version](#downloading-and-setting-your-default-nodejs-to-a-new-version)
        1. [Examples](#examples-3)
    7. [Printing your current version of Node.js](#printing-your-current-version-of-nodejs)
    8. [Printing all downloaded versions of Node.js](#printing-all-downloaded-versions-of-nodejs)
    9. [Deleting a downloaded version of Node.js](#deleting-a-downloaded-version-of-nodejs)
        1. [Examples](#examples-4)
    10. [Printing your current version of PolyNode](#printing-your-current-version-of-polynode)
5. [How to configure](#how-to-configure-polynode)
    1. [Configuration fields](#configuration-fields)
        1. [nodeMirror](#nodemirror)
6. [How to uninstall](#how-to-uninstall-polynode)
    1. [AIX, Linux, or macOS](#aix-linux-or-macos)
    2. [Windows](#windows)
7. [Building from source](#building-from-source)
    1. [Bundle scripts](#bundle-scripts)
        1. [Required technologies](#required-technologies)
        2. [AIX](#building-on-aix)
        3. [Linux](#building-on-linux)
        4. [macOS](#building-on-macos)
        5. [Windows](#building-on-windows)
    2. [Dockerfile](#dockerfile)
        1. [Required technologies](#required-technologies-1)
        2. [Building an image](#building-an-image)
8. [Contributing](#contributing)
9. [Acknowledgements](#acknowledgements)

## Quick demo

![polyn demo](https://github.com/sionpixley/PolyNode/blob/main/demo.gif)

## Supported operating systems and CPU architectures

- AIX (Power 64-bit)
- Linux (x64, ARM64, Power LE 64-bit, and s390x)
- macOS (x64 and ARM64)
- Windows (x64 and ARM64)

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
2. Download the Linux .tar.xz or .tar.gz file appropriate for your CPU architecture.
3. Extract the tarball and run the setup binary.

### For macOS

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the Darwin .tar.gz file appropriate for your CPU architecture.
3. Extract the .tar.gz file and run the setup binary.

### For Windows

1. Navigate to the [latest release](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the Windows .zip file appropriate for your CPU architecture.
3. Extract the .zip file and run setup.exe.

## How to use

PolyNode does not require sudo/admin privileges to use the `polyn` command.

### Updating PolyNode to the latest release

PolyNode has an auto updater, so manually updating your PolyNode is not usually required.

You can turn off the auto updater in the [configuration file](#how-to-configure-polynode).

`polyn update`

### Searching for available Node.js versions

`polyn search [prefix]`

Using a prefix will match anything with this prefix. So `polyn search 1` will match with any Node.js version that starts with "1".

If you want to search for a specific major version, add a "." at the end of your prefix. `polyn search 18.` will print all Node.js v18 releases.

A default list will print if no prefix is given.

### Downloading a new version of Node.js

This command will only download a new version of Node.js. It will not set the new version as your currently-used version. See [Setting your default Node.js version](#setting-your-default-nodejs-version) or [Temporarily setting your Node.js version](#temporarily-setting-your-nodejs-version) on how to use the Node.js you download.

`polyn add <version or keyword or prefix>`

#### Examples

```sh
# Downloading a specific version of Node.js.
polyn add 23.7.0

# Downloading the latest Node.js release that matches a prefix.
polyn add 23

# Downloading the latest Node.js LTS release.
polyn add lts

# Downloading the latest Node.js release.
polyn add latest
```

### Setting your default Node.js version

This command will set your Node.js version across all shell processes. All new shell processes will automatically use this Node.js version, unless overriden by [temporarily setting the Node.js version](#temporarily-setting-your-nodejs-version).

`polyn use <version or prefix>`

#### Examples

```sh
# Setting your default to a specific Node.js version.
polyn use 23.7.0

# Setting your default to the latest Node.js release that matches a prefix.
polyn use 23
```

### Temporarily setting your Node.js version

This command will temporarily set your Node.js version for your current shell process and all child processes of that shell. This will only set your Node.js version for the lifetime of the shell. For a more permanent solution, see [Setting your default Node.js version](#setting-your-default-nodejs-version).

This command is useful if you need to run two separate projects at the same time that depend on different versions of Node.js. 

#### Temporarily setting your Node.js on AIX, Linux, or macOS

`eval $(polyn temp <version or prefix>)`

##### Examples

```sh
# Temporarily setting your Node.js to a specific version.
eval $(polyn temp 23.7.0)

# Temporarily setting your Node.js to the latest release that matches a prefix.
eval $(polyn temp 23)
```

#### Temporarily setting your Node.js on Windows

Unfortunately, Windows doesn't have a command equivalent to the POSIX `eval`. You will have to run `polyn temp <version or prefix>` and then copy and paste the command it outputs.

### Downloading and setting your default Node.js to a new version

This command downloads a specific version of Node.js and immediately sets it as your default version.

The `install` command is equivalent to the `add` command followed by the `use` command.

`polyn install <version or keyword or prefix>`

#### Examples

```sh
# Downloading and setting your default to a specific version of Node.js.
polyn install 23.7.0

# Downloading and setting your default to the latest Node.js release that matches a prefix.
polyn install 23

# Downloading and setting your default to the latest Node.js LTS release.
polyn install lts

# Downloading and setting your default to the latest Node.js release.
polyn install latest
```

### Printing your current version of Node.js

`polyn current`

### Printing all downloaded versions of Node.js

`polyn ls`

or 

`polyn list`

### Deleting a downloaded version of Node.js

`polyn rm <version or prefix>`

or 

`polyn remove <version or prefix>`

#### Examples

```sh
# Deleting a specific version of Node.js.
polyn rm 23.7.0

# Deleting the oldest Node.js release that matches a prefix.
polyn rm 23
```

### Printing your current version of PolyNode

`polyn version`

## How to configure PolyNode

PolyNode's configuration is handled through a JSON file named `polynrc.json` located in PolyNode's home directory (`$HOME/.PolyNode` for AIX/Linux/macOS or `%LOCALAPPDATA%\Programs\PolyNode` for Windows). Please see below for the default configuration for `polynrc.json`:

```json
{
  "autoUpdate": true,
  "nodeMirror": "https://nodejs.org/dist"
}
```

### Configuration fields

#### autoUpdate

This field is a `bool` that configures if PolyNode's auto updater should run. Default value is `true`.

#### nodeMirror

This field is a `string` that configures the URL to download Node.js. Default value is `"https://nodejs.org/dist"`.

## How to uninstall PolyNode

PolyNode does not require sudo/admin privileges to uninstall.

### AIX, Linux, or macOS

1. Run the `$HOME/.PolyNode/uninstall/uninstall` binary.

### Windows

1. Run `%LOCALAPPDATA%\Programs\PolyNode\uninstall\uninstall.exe`.

## Building from source

There are two main ways to build PolyNode from source: Using the [bundle scripts](#bundle-scripts) or [building the Dockerfile](#dockerfile). 

If you're just testing your build locally, I would recommend building a Docker image from the Dockerfile. The bundle scripts are helpful if you want to install/distribute your own build.

### Bundle scripts

#### Required technologies

- Go 1.25.3

#### Building on AIX

Run the POSIX shell script `./scripts/aix/bundle`. This script will build PolyNode's source code for Power 64-bit and bundle the artifacts as a .tar.gz file.

#### Building on Linux

Run the POSIX shell script `./scripts/linux/bundle`. This script will build PolyNode's source code for x64, ARM64, Power LE 64-bit, and s390x and bundle the artifacts as separate .tar.xz and .tar.gz files. The contents of the .tar.xz files and the .tar.gz files are identical. Both formats are provided for backwards compatiblity reasons.

#### Building on macOS

macOS has a POSIX shell script (`./scripts/mac/bundle`) that builds and notarizes PolyNode's source code for x64 and ARM64 and bundles the artifacts as separate .tar.gz files. If you don't need to distribute the binaries, then you don't need the notarization step. Just edit the bundle script and set the `sign` variable to `0`.

#### Building on Windows

Run the batchfile `.\scripts\win\bundle.cmd`. This batchfile will build PolyNode's source code for x64 and ARM64 and bundle the artifacts as separate .zip files.

### Dockerfile

#### Required technologies

- Docker

#### Building an image

`docker build -t polyn .`

## Contributing

All contributions are welcome! If you wish to contribute to the project, the best way would be forking this repo and making a pull request from your fork with all of your suggested changes.

## Acknowledgements

PolyNode draws a lot of inspiration, especially in regards to syntax, from other, more well-known projects, like: [nvm](https://github.com/nvm-sh/nvm), [nvm-windows](https://github.com/coreybutler/nvm-windows), and [nvs](https://github.com/jasongin/nvs).

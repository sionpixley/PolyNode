# PolyNode

PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device. It also comes with a GUI app that you can use instead.

## Table of contents

1. [Supported operating systems and CPU architectures](#supported-operating-systems-and-cpu-architectures)
    1. [Linux support information](#linux-support-information)
2. [How to install](#how-to-install-polynode)
    1. [Linux](#for-linux)
    2. [macOS](#for-macos)
    3. [Windows](#for-windows)
3. [How to use](#how-to-use)
    1. [Searching for available Node.js versions](#searching-for-available-nodejs-versions)
    2. [Searching for a specific Node.js version](#searching-for-a-specific-nodejs-version)
    3. [Downloading and switching to a new version of Node.js](#downloading-and-switching-to-a-new-version-of-nodejs)
    4. [Downloading a new version of Node.js](#downloading-a-new-version-of-nodejs)
    5. [Switching to a different downloaded version of Node.js](#switching-to-a-different-downloaded-version-of-nodejs)
    6. [Printing your current version of Node.js](#printing-your-current-version-of-nodejs)
    7. [Printing all downloaded versions of Node.js](#printing-all-downloaded-versions-of-nodejs)
    8. [Deleting a downloaded version of Node.js](#deleting-a-downloaded-version-of-nodejs)
    9. [Printing your current version of PolyNode](#printing-your-current-version-of-polynode)
    10. [Launching the GUI](#launching-the-gui)
4. [How to configure](#how-to-configure-polynode)
    1. [Configuration fields](#configuration-fields)
5. [How to uninstall](#how-to-uninstall-polynode)
    1. [Linux and macOS](#linux-and-macos)
    2. [Windows](#windows)
6. [Future development](#future-development)
7. [Information](#information)

## Supported operating systems and CPU architectures

- Linux (x64 and ARM64)
- macOS (x64 and ARM64)
- Windows 10 and newer (x64 and ARM64)

### Linux support information

PolyNode only supports Bash or Zsh by default. During the install process, PolyNode edits either .bashrc or .zshrc to add two locations to the PATH: PolyNode's home directory `~/.PolyNode` and the symlink for Node.js `~/.PolyNode/nodejs`. You can get PolyNode to work for other shells by adding these directories to your PATH environment variable.

## How to install PolyNode

PolyNode does not require sudo/admin privileges to install.

If you have a previous version of PolyNode installed, you do not have to uninstall it before installing the new version.

Please uninstall all Node.js downloads that weren't installed by PolyNode before running the setup executable.

### For Linux

1. Navigate to [Releases](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the latest Linux .tar.xz file appropriate for your CPU architecture.
3. Extract the .tar.xz file and run the setup executable.

### For macOS

1. Navigate to [Releases](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the latest Darwin .tar.gz file appropriate for your CPU architecture.
3. Extract the .tar.gz file and run the setup executable.

### For Windows

1. Navigate to [Releases](https://github.com/sionpixley/PolyNode/releases/latest).
2. Download the latest Windows .zip file appropriate for your CPU architecture.
3. Extract the .zip file and run setup.exe.

## How to use

PolyNode does not require sudo/admin privileges to use the `polyn` command.

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

### Launching the GUI

If you'd rather work with a GUI instead of using the CLI, type this command into your terminal:

`PolyNode`

## How to configure PolyNode

PolyNode's configuration is handled through a JSON file named `.polynrc` located in PolyNode's home directory (`~/.PolyNode` for Linux/macOS and `%LOCALAPPDATA%\Programs\PolyNode` for Windows). Please see below for the default configuration for `.polynrc`:

```
{
  "nodeMirror": "https://nodejs.org/dist"
}
```

This configuration file is limited at the moment. I hope to expand its capabilities over time.

### Configuration fields

#### nodeMirror

This field is a `string` that represents the URL to download Node.js. Default value is `"https://nodejs.org/dist"`.

## How to uninstall PolyNode

PolyNode does not require sudo/admin privileges to uninstall.

### Linux and macOS

1. Run the `~/.PolyNode/uninstall/uninstall` executable.

### Windows

1. Run `%LOCALAPPDATA%\Programs\PolyNode\uninstall\uninstall.exe`.

## Future development

The original scope of this project was to be able to install and manage multiple versions of Bun, Deno, and Node.js. It currently only supports Node.js, but I would like to support Bun and Deno in the future.

## Information

PolyNode draws a lot of inspiration, especially in regards to syntax, from other, more well-known projects, like: [nvm](https://github.com/nvm-sh/nvm), [nvm-windows](https://github.com/coreybutler/nvm-windows), and [nvs](https://github.com/jasongin/nvs).

Go 1.23
Angular 18
pnpm 9.10.0

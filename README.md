# PolyNode

PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device. Primarily made as a side-project to help me learn Go.

## Supported operating systems and CPU architectures

- Linux ARM64
- Linux x64
- macOS ARM64
- macOS x64
- Windows x64

### Linux support

PolyNode only supports Bash or Zsh by default. During the install process, PolyNode edits either .bashrc or .zshrc to add two locations to the PATH: PolyNode's home directory `~/.PolyNode` and the symlink for Node.js and related commands `~/.PolyNode/nodejs`. This will probably change later to have support for more shells.

## How to install

Please uninstall all installed versions of Node.js before installing PolyNode.

### For Linux

#### ARM64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/PolyNode/releases) section and download the linux-\<version\>-arm64.tar.xz file.
2. Extract the .tar.xz file and run the `setup` executable.

#### x86-64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/PolyNode/releases) section and download the linux-\<version\>-x64.tar.xz file.
2. Extract the .tar.xz file and run the `setup` executable.

### For macOS

#### ARM64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/PolyNode/releases) section and download the darwin-\<version\>-arm64.tar.gz file.
2. Extract the .tar.gz file and run the `setup` executable.

#### x86-64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/PolyNode/releases) section and download the darwin-\<version\>-x64.tar.gz file.
2. Extract the .tar.gz file and run the `setup` executable.

### For Windows

1. Navigate to the [Releases](https://github.com/sionpixley/PolyNode/releases) section and download the win-\<version\>-x64.zip file.
2. Extract the .zip file and run the `setup.exe` executable.

## How to use

### Downloading a new version of Node.js

`polyn add <version>`

### Switching to a different downloaded version of Node.js

`polyn use <version>`

### Downloading and switching to a new version of Node.js

`polyn install <version>`

### Printing the current version of Node.js

`polyn current`

### Printing all downloaded versions of Node.js

`polyn ls`

or 

`polyn list`

### Deleting a downloaded version of Node.js

`polyn rm <version>`

or 

`polyn remove <version>`

## Future development

The original scope of this project was to be able to install and manage multiple versions of Bun, Deno, and Node.js. It currently only supports Node.js, but I would like to support Bun and Deno in the future.

PolyNode currently cannot search for available versions of Node.js. I would like to add that soon.

## Information

Go v1.22.4 <br>
7-Zip 24.05

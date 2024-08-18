# PolyNode

PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device.

## Supported operating systems and CPU architectures

- Linux ARM64
- Linux x64
- macOS ARM64
- macOS x64

### Linux support

PolyNode only supports Bash or Zsh by default. During the install process, PolyNode edits either .bashrc or .zshrc to add two locations to the PATH: PolyNode's home directory `~/.PolyNode` and the symlink for Node.js `~/.PolyNode/nodejs`. This will probably change later to have support for more shells.

## How to install PolyNode

PolyNode does not require sudo privileges to install. Please uninstall all installed versions of Node.js before installing PolyNode.

### For Linux

1. Navigate to [Releases](https://github.com/sionpixley/PolyNode/releases).
2. Download the latest Linux .tar.xz file appropriate for your CPU architecture.
3. Extract the .tar.xz file and run the setup executable.

### For macOS

1. Navigate to [Releases](https://github.com/sionpixley/PolyNode/releases).
2. Download the latest Darwin .tar.gz file appropriate for your CPU architecture.
3. Extract the .tar.gz file and run the setup executable.

## How to use

PolyNode does not require sudo privileges to use the `polyn` command.

### Printing the most recent Node.js versions available for download

`polyn search`

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

## How to uninstall PolyNode

PolyNode does not require sudo privileges to uninstall.

### Linux and macOS

1. Run the `~/.PolyNode/uninstall/uninstall` executable.

## Future development

The original scope of this project was to be able to install and manage multiple versions of Bun, Deno, and Node.js. It currently only supports Node.js, but I would like to support Bun and Deno in the future.

It also doesn't currently support Windows. I'll fix that later.

PolyNode currently cannot search for available versions of Node.js. I would like to add that soon.

## Information

Go 1.22

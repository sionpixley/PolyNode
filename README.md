# PolyNode

PolyNode is a CLI tool that helps install and manage multiple versions of Node.js on the same device. Primarily made as a side project to help me learn Go. I do not recommend using it in production in its current state.

## Supported operating systems and CPU architectures

- Linux ARM64
- Linux x64
- macOS ARM64
- macOS x64
- Windows x64

### Linux support

PolyNode only supports Bash or Zsh by default. During the install process, PolyNode edits either .bashrc or .zshrc to add two locations to the PATH: PolyNode's home directory `~/.PolyNode` and the symlink for Node.js `~/.PolyNode/nodejs`. This will probably change later to have support for more shells.

### Windows support

The uninstall executable isn't completely finished yet. It does remove PolyNode, but does not remove it from the PATH.

Additionally, on Windows, PolyNode currently requires admin privileges to use the `polyn` command. This is because PolyNode creates a symlink to the directory that it stores Node.js commands/packages (e.g. node, npm). Symlinks require admin rights on Windows. The original idea was for PolyNode to not require any sudo/admin privileges at all. I'll figure something out.

## How to install PolyNode

PolyNode does not require sudo/admin privileges to install.

Please uninstall all installed versions of Node.js before installing PolyNode. Please make sure Go 1.22 is already installed on your machine before proceeding with the following steps.

### For Linux

1. Clone or download this repo.
2. In your terminal, run the bundle script located at `scripts/linux/bundle.sh`.
3. The bundle script will create two new directories: `linux-<version>-arm64` and `linux-<version>-x64`. Run the `setup` executable found in the correct directory for your CPU architecture.

### For macOS

1. Clone or download this repo.
2. In your terminal, run the bundle script located at `scripts/mac/bundle.zsh`.
3. The bundle script will create two new directories: `darwin-<version>-arm64` and `darwin-<version>-x64`. Run the `setup` executable found in the correct directory for your CPU architecture.

### For Windows

1. Clone or download this repo.
2. In your terminal, run the bundle batchfile located at `scripts/win/bundle.cmd`.
3. The bundle batchfile will create a new directory: `win-<version>-x64`. Run `setup.exe` found in that directory.

## How to use

On Linux and macOS, PolyNode does not require sudo privileges to use the `polyn` command.

On Windows, PolyNode **does** require admin privileges to use the `polyn` command. I will be looking to change this. For now, please launch cmd/PowerShell/Terminal as admin before using the `polyn` command.

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

PolyNode does not require any sudo/admin privileges to uninstall.

### Linux and macOS

1. Run the `~/.PolyNode/uninstall/uninstall` executable.

### Windows

1. Run `%LOCALAPPDATA%\Programs\PolyNode\uninstall\uninstall.exe`.
2. Remove `%LOCALAPPDATA%\Programs\PolyNode` from the user PATH.
3. Remove `%LOCALAPPDATA%\Programs\PolyNode\nodejs` from the user PATH.

## Future development

The original scope of this project was to be able to install and manage multiple versions of Bun, Deno, and Node.js. It currently only supports Node.js, but I would like to support Bun and Deno in the future.

On Windows, the `polyn` command requires admin privileges to make a symlink. I would like to figure out a way around this so that admin isn't required.

PolyNode currently cannot search for available versions of Node.js. I would like to add that soon.

## Information

Go 1.22

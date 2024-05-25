# PolyNode

PolyNode is a CLI tool that helps install and switch between multiple versions of Bun, Deno, and/or Node.js on the same device.

## Supported operating systems and CPU architectures

- Linux ARM64
- Linux x64
- macOS ARM64
- macOS x64
- Windows x64

## How to install

### For Linux

#### ARM64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/polyn/releases) section and download the linux-\<version\>-arm64.tar.xz file
2. Extract the .tar.xz file and run the `setup` executable

#### x86-64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/polyn/releases) section and download the linux-\<version\>-x64.tar.xz file
2. Extract the .tar.xz file and run the `setup` executable

### For macOS

#### ARM64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/polyn/releases) section and download the darwin-\<version\>-arm64.tar.gz file
2. Extract the .tar.gz file and run the `setup` executable

#### x86-64 based CPUs

1. Navigate to the [Releases](https://github.com/sionpixley/polyn/releases) section and download the darwin-\<version\>-x64.tar.gz file
2. Extract the .tar.gz file and run the `setup` executable

### For Windows

1. Navigate to the [Releases](https://github.com/sionpixley/polyn/releases) section and download the win-\<version\>-x64.zip file
2. Extract the .zip file and run the `setup.exe` executable

## How to use

### Downloading a new version of a runtime

#### Bun

`polyn bun add <version>`

#### Deno

`polyn deno add <version>`

#### Node.js

`polyn add <version>`

or 

`polyn node add <version>`

### Switching to a different downloaded version of a runtime

#### Bun

`polyn bun use <version>`

#### Deno

`polyn deno use <version>`

#### Node.js

`polyn use <version>`

or

`polyn node use <version>`

### Downloading and switching to a new version of a runtime

#### Bun

`polyn bun install <version>`

#### Deno

`polyn deno install <version>`

#### Node.js

`polyn install <version>`

or 

`polyn node install <version>`

### Printing the current version of a runtime

#### Bun

`polyn bun current`

#### Deno

`polyn deno current`

#### Node.js

`polyn current`

or 

`polyn node current`

### Printing downloaded versions of a runtime

#### Bun

`polyn bun ls`

or 

`polyn bun list`

#### Deno

`polyn deno ls`

or 

`polyn deno list`

#### Node.js

`polyn ls`

or 

`polyn node ls`

or 

`polyn list`

or

`polyn node list`

### Deleting a downloaded version of a runtime

#### Bun

`polyn bun rm <version>`

or 

`polyn bun remove <version>`

#### Deno

`polyn deno rm <version>`

or 

`polyn deno remove <version>`

#### Node.js

`polyn rm <version>`

or 

`polyn node rm <version>`

or 

`polyn remove <version>`

or

`polyn node remove <version>`

## Information

Go v1.22.3 <br>
7-Zip 24.05

#!/bin/zsh

version=v0.1.0

# Build macOS ARM64
env GOOS=darwin GOARCH=arm64 go build -tags=prod -o polyn-darwin-arm64 ./cmd/polyn

cd install
env GOOS=darwin GOARCH=arm64 go build -o ../setup-darwin-arm64

cd ../uninstall
env GOOS=darwin GOARCH=arm64 go build -o ../uninstall-darwin-arm64

cd ..
mkdir PolyNode-$version-darwin-arm64
mv setup-darwin-arm64 ./PolyNode-$version-darwin-arm64/setup

cd PolyNode-$version-darwin-arm64
mkdir PolyNode
cd ..
mv polyn-darwin-arm64 ./PolyNode-$version-darwin-arm64/PolyNode/polyn

cp README.md ./PolyNode-$version-darwin-arm64/PolyNode
cp LICENSE ./PolyNode-$version-darwin-arm64/PolyNode

cd ./PolyNode-$version-darwin-arm64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-darwin-arm64 ./PolyNode-$version-darwin-arm64/PolyNode/uninstall/uninstall

# Build macOS x64
env GOOS=darwin GOARCH=amd64 go build -tags=prod -o polyn-darwin-x64 ./cmd/polyn

cd install
env GOOS=darwin GOARCH=amd64 go build -o ../setup-darwin-x64

cd ../uninstall
env GOOS=darwin GOARCH=amd64 go build -o ../uninstall-darwin-x64

cd ..
mkdir PolyNode-$version-darwin-x64
mv setup-darwin-x64 ./PolyNode-$version-darwin-x64/setup

cd PolyNode-$version-darwin-x64
mkdir PolyNode
cd ..
mv polyn-darwin-x64 ./PolyNode-$version-darwin-x64/PolyNode/polyn

cp README.md ./PolyNode-$version-darwin-x64/PolyNode
cp LICENSE ./PolyNode-$version-darwin-x64/PolyNode

cd ./PolyNode-$version-darwin-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-darwin-x64 ./PolyNode-$version-darwin-x64/PolyNode/uninstall/uninstall

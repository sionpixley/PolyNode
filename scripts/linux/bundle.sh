#!/bin/bash

version=v0.6.1

# Build Linux ARM64

env GOOS=linux GOARCH=arm64 go build -tags=prod -o polyn-linux-arm64 ./cmd/polyn

cd install
env GOOS=linux GOARCH=arm64 go build -o ../setup-linux-arm64

cd ../uninstall
env GOOS=linux GOARCH=arm64 go build -o ../uninstall-linux-arm64

cd ..
mkdir PolyNode-$version-linux-arm64
mv setup-linux-arm64 ./PolyNode-$version-linux-arm64/setup

cd PolyNode-$version-linux-arm64
mkdir PolyNode
cd ..
mv polyn-linux-arm64 ./PolyNode-$version-linux-arm64/PolyNode/polyn

cp README.md ./PolyNode-$version-linux-arm64/PolyNode
cp LICENSE ./PolyNode-$version-linux-arm64/PolyNode
cp SECURITY.md ./PolyNode-$version-linux-arm64/PolyNode

cd ./PolyNode-$version-linux-arm64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-arm64 ./PolyNode-$version-linux-arm64/PolyNode/uninstall/uninstall

# Build Linux x64

env GOOS=linux GOARCH=amd64 go build -tags=prod -o polyn-linux-x64 ./cmd/polyn

cd install
env GOOS=linux GOARCH=amd64 go build -o ../setup-linux-x64

cd ../uninstall
env GOOS=linux GOARCH=amd64 go build -o ../uninstall-linux-x64

cd ..
mkdir PolyNode-$version-linux-x64
mv setup-linux-x64 ./PolyNode-$version-linux-x64/setup

cd PolyNode-$version-linux-x64
mkdir PolyNode
cd ..
mv polyn-linux-x64 ./PolyNode-$version-linux-x64/PolyNode/polyn

cp README.md ./PolyNode-$version-linux-x64/PolyNode
cp LICENSE ./PolyNode-$version-linux-x64/PolyNode
cp SECURITY.md ./PolyNode-$version-linux-x64/PolyNode

cd ./PolyNode-$version-linux-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-x64 ./PolyNode-$version-linux-x64/PolyNode/uninstall/uninstall

# Bundle

tar -cJf PolyNode-$version-linux-arm64.tar.xz PolyNode-$version-linux-arm64
rm -rf PolyNode-$version-linux-arm64

tar -cJf PolyNode-$version-linux-x64.tar.xz PolyNode-$version-linux-x64
rm -rf PolyNode-$version-linux-x64

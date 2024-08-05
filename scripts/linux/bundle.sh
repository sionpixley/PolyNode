#!/bin/bash

version=v0.1.0

# Build Linux ARM64
env GOOS=linux GOARCH=arm64 go build -tags=prod -o polyn-linux-arm64 ./cmd/polyn

cd install
env GOOS=linux GOARCH=arm64 go build -o ../setup-linux-arm64

cd ../uninstall
env GOOS=linux GOARCH=arm64 go build -o ../uninstall-linux-arm64

cd ..
mkdir linux-$version-arm64
mv setup-linux-arm64 ./linux-$version-arm64/setup

cd linux-$version-arm64
mkdir PolyNode
cd ..
mv polyn-linux-arm64 ./linux-$version-arm64/PolyNode/polyn

cp README.md ./linux-$version-arm64/PolyNode
cp LICENSE ./linux-$version-arm64/PolyNode

cd ./linux-$version-arm64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-arm64 ./linux-$version-arm64/PolyNode/uninstall/uninstall

# Build Linux x64
env GOOS=linux GOARCH=amd64 go build -tags=prod -o polyn-linux-x64 ./cmd/polyn

cd install
env GOOS=linux GOARCH=amd64 go build -o ../setup-linux-x64

cd ../uninstall
env GOOS=linux GOARCH=amd64 go build -o ../uninstall-linux-x64

cd ..
mkdir linux-$version-x64
mv setup-linux-x64 ./linux-$version-x64/setup

cd linux-$version-x64
mkdir PolyNode
cd ..
mv polyn-linux-x64 ./linux-$version-x64/PolyNode/polyn

cp README.md ./linux-$version-x64/PolyNode
cp LICENSE ./linux-$version-x64/PolyNode

cd ./linux-$version-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-x64 ./linux-$version-x64/PolyNode/uninstall/uninstall

#!/bin/bash

version=v0.1.0

# Build Linux ARM64
env GOOS=linux GOARCH=arm64 go build -tags=prod -o polyn-linux-arm64 ./cmd/polyn

cd install
go build -o ../setup-linux-arm64

cd ../uninstall
go build -o ../uninstall-linux-arm64

cd ..
mkdir linux-$version-arm64
mv setup-linux-arm64 ./linux-$version-arm64/setup

cd linux-$version-arm64
mkdir PolyNode
cd ..
mv polyn-linux-arm64 ./linux-$version-arm64/PolyNode/polyn

cp README.md ./linux-$version-arm64/PolyNode
cp LICENSE.md ./linux-$version-arm64/PolyNode

cd ./linux-$version-arm64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-arm64 ./linux-$version-arm64/PolyNode/uninstall/uninstall

cd ./linux-$version-arm64/PolyNode
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir linux
cd linux
mkdir arm64
cd ../../../../..

cp ./emb/7z/linux/arm64/7zzs ./linux-$version-arm64/PolyNode/emb/7z/linux/arm64
cp ./emb/7z/linux/License.txt ./linux-$version-arm64/PolyNode/emb/7z/linux

# Build Linux x64
env GOOS=linux GOARCH=amd64 go build -tags=prod -o polyn-linux-x64 ./cmd/polyn

cd install
go build -o ../setup-linux-x64

cd ../uninstall
go build -o ../uninstall-linux-x64

cd ..
mkdir linux-$version-x64
mv setup-linux-x64 ./linux-$version-x64/setup

cd linux-$version-x64
mkdir PolyNode
cd ..
mv polyn-linux-x64 ./linux-$version-x64/PolyNode/polyn

cp README.md ./linux-$version-x64/PolyNode
cp LICENSE.md ./linux-$version-x64/PolyNode

cd ./linux-$version-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-linux-x64 ./linux-$version-x64/PolyNode/uninstall/uninstall

cd ./linux-$version-x64/PolyNode
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir linux
cd linux
mkdir x64
cd ../../../../..

cp ./emb/7z/linux/x64/7zzs ./linux-$version-x64/PolyNode/emb/7z/linux/x64
cp ./emb/7z/linux/License.txt ./linux-$version-x64/PolyNode/emb/7z/linux

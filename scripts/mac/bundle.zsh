#!/bin/zsh

version=v0.1.0

# Build macOS ARM64
env GOOS=darwin GOARCH=arm64 go build -tags=prod -o polyn-darwin-arm64 ./cmd/polyn

cd install
go build -o ../setup-darwin-arm64

cd ../uninstall
go build -o ../uninstall-darwin-arm64

cd ..
mkdir darwin-$version-arm64
mv setup-darwin-arm64 ./darwin-$version-arm64/setup

cd darwin-$version-arm64
mkdir PolyNode
cd ..
mv polyn-darwin-arm64 ./darwin-$version-arm64/PolyNode/polyn

cp README.md ./darwin-$version-arm64/PolyNode
cp LICENSE.md ./darwin-$version-arm64/PolyNode

cd ./darwin-$version-arm64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-darwin-arm64 ./darwin-$version-arm64/PolyNode/uninstall/uninstall

cd ./darwin-$version-arm64/PolyNode
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir mac
cd ../../../..

cp ./emb/7z/mac/7zz ./darwin-$version-arm64/PolyNode/emb/7z/mac
cp ./emb/7z/mac/License.txt ./darwin-$version-arm64/PolyNode/emb/7z/mac

# Build macOS x64
env GOOS=darwin GOARCH=amd64 go build -tags=prod -o polyn-darwin-x64 ./cmd/polyn

cd install
go build -o ../setup-darwin-x64

cd ../uninstall
go build -o ../uninstall-darwin-x64

cd ..
mkdir darwin-$version-x64
mv setup-darwin-x64 ./darwin-$version-x64/setup

cd darwin-$version-x64
mkdir PolyNode
cd ..
mv polyn-darwin-x64 ./darwin-$version-x64/PolyNode/polyn

cp README.md ./darwin-$version-x64/PolyNode
cp LICENSE.md ./darwin-$version-x64/PolyNode

cd ./darwin-$version-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-darwin-x64 ./darwin-$version-x64/PolyNode/uninstall/uninstall

cd ./darwin-$version-x64/PolyNode
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir mac
cd ../../../..

cp ./emb/7z/mac/7zz ./darwin-$version-x64/PolyNode/emb/7z/mac
cp ./emb/7z/mac/License.txt ./darwin-$version-x64/PolyNode/emb/7z/mac

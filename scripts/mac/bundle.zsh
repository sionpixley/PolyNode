#!/bin/zsh

version=v0.3.0

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
cp SECURITY.md ./PolyNode-$version-darwin-arm64/PolyNode
cp CODE_OF_CONDUCT.md ./PolyNode-$version-darwin-arm64/PolyNode

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
cp SECURITY.md ./PolyNode-$version-darwin-x64/PolyNode
cp CODE_OF_CONDUCT.md ./PolyNode-$version-darwin-x64/PolyNode

cd ./PolyNode-$version-darwin-x64/PolyNode
mkdir uninstall
cd ../..
mv uninstall-darwin-x64 ./PolyNode-$version-darwin-x64/PolyNode/uninstall/uninstall

# Notarize

cd PolyNode-$version-darwin-arm64
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID setup
zip setup.zip setup
rm setup
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD setup.zip
unzip setup.zip
rm setup.zip

cd PolyNode
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID polyn
zip polyn.zip polyn
rm polyn
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD polyn.zip
unzip polyn.zip
rm polyn.zip

cd uninstall
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID uninstall
zip uninstall.zip uninstall
rm uninstall
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD uninstall.zip
unzip uninstall.zip
rm uninstall.zip
cd ../../..

cd PolyNode-$version-darwin-x64
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID setup
zip setup.zip setup
rm setup
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD setup.zip
unzip setup.zip
rm setup.zip

cd PolyNode
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID polyn
zip polyn.zip polyn
rm polyn
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD polyn.zip
unzip polyn.zip
rm polyn.zip

cd uninstall
codesign -v -f --timestamp --options=runtime --sign $CODESIGN_TEAM_ID uninstall
zip uninstall.zip uninstall
rm uninstall
xcrun notarytool submit --apple-id $CODESIGN_APPLE_ID --team-id $CODESIGN_TEAM_ID --password $CODESIGN_PASSWD uninstall.zip
unzip uninstall.zip
rm uninstall.zip
cd ../../..

# Bundle

tar -czf PolyNode-$version-darwin-arm64.tar.gz PolyNode-$version-darwin-arm64
rm -rf PolyNode-$version-darwin-arm64

tar -czf PolyNode-$version-darwin-x64.tar.gz PolyNode-$version-darwin-x64
rm -rf PolyNode-$version-darwin-x64

@echo off

:: Build Windows ARM64

set GOOS=windows
set GOARCH=arm64

go build -tags=prod -o polyn.exe ./cmd/polyn

cd install
go build -o ../setup.exe ./cmd/setup

cd ..\uninstall
go build -o ../uninstall-arm64.exe

cd ..
mkdir PolyNode-win-arm64
move setup.exe .\PolyNode-win-arm64\setup.exe

cd PolyNode-win-arm64
mkdir PolyNode
cd ..
copy polyn.exe .\PolyNode-win-arm64\PolyNode

copy README.md .\PolyNode-win-arm64\PolyNode
copy LICENSE .\PolyNode-win-arm64\PolyNode
copy SECURITY.md .\PolyNode-win-arm64\PolyNode

cd .\PolyNode-win-arm64\PolyNode
mkdir uninstall
cd ..\..
move uninstall-arm64.exe .\PolyNode-win-arm64\PolyNode\uninstall\uninstall.exe

:: Build Windows x64

set GOOS=windows
set GOARCH=amd64

go build -tags=prod -o polyn.exe ./cmd/polyn

cd install
go build -o ../setup.exe ./cmd/setup

cd ..\uninstall
go build -o ../uninstall-x64.exe

cd ..
mkdir PolyNode-win-x64
move setup.exe .\PolyNode-win-x64\setup.exe

cd PolyNode-win-x64
mkdir PolyNode
cd ..
copy polyn.exe .\PolyNode-win-x64\PolyNode

copy README.md .\PolyNode-win-x64\PolyNode
copy LICENSE .\PolyNode-win-x64\PolyNode
copy SECURITY.md .\PolyNode-win-x64\PolyNode

cd .\PolyNode-win-x64\PolyNode
mkdir uninstall
cd ..\..
move uninstall-x64.exe .\PolyNode-win-x64\PolyNode\uninstall\uninstall.exe

:: Bundle

powershell -Command "Compress-Archive -Path PolyNode-win-arm64 -DestinationPath PolyNode-win-arm64.zip"
del PolyNode-win-arm64 /s /f /q > nul
rmdir PolyNode-win-arm64 /s /q

powershell -Command "Compress-Archive -Path PolyNode-win-x64 -DestinationPath PolyNode-win-x64.zip"
del PolyNode-win-x64 /s /f /q > nul
rmdir PolyNode-win-x64 /s /q

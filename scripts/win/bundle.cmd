@echo off

set version=v0.8.0

:: Build Windows ARM64

set GOOS=windows
set GOARCH=arm64

go build -tags=prod -o polyn-arm64.exe ./cmd/polyn

cd install
go build -o ../setup-arm64.exe ./cmd/setup

cd ..\uninstall
go build -o ../uninstall-arm64.exe

cd ..\web
go build -tags=prod -o ../PolyNode-arm64.exe ./cmd/serve

cd gui
pnpm install
pnpm run build

cd ..\..
mkdir PolyNode-%version%-win-arm64
move setup-arm64.exe .\PolyNode-%version%-win-arm64\setup.exe

cd PolyNode-%version%-win-arm64
mkdir PolyNode
cd ..
move polyn-arm64.exe .\PolyNode-%version%-win-arm64\PolyNode\polyn.exe
move PolyNode-arm64.exe .\PolyNode-%version%-win-arm64\PolyNode\PolyNode.exe

copy README.md .\PolyNode-%version%-win-arm64\PolyNode
copy LICENSE .\PolyNode-%version%-win-arm64\PolyNode
copy SECURITY.md .\PolyNode-%version%-win-arm64\PolyNode

cd .\PolyNode-%version%-win-arm64\PolyNode
mkdir uninstall
mkdir gui
cd ..\..
move uninstall-arm64.exe .\PolyNode-%version%-win-arm64\PolyNode\uninstall\uninstall.exe
xcopy /s /i .\web\gui\dist\ .\PolyNode-%version%-win-arm64\PolyNode\gui\dist\

:: Build Windows x64

set GOOS=windows
set GOARCH=amd64

go build -tags=prod -o polyn-x64.exe ./cmd/polyn

cd install
go build -o ../setup-x64.exe ./cmd/setup

cd ..\uninstall
go build -o ../uninstall-x64.exe

cd ..\web
go build -tags=prod -o ../PolyNode-x64.exe ./cmd/serve

cd ..
mkdir PolyNode-%version%-win-x64
move setup-x64.exe .\PolyNode-%version%-win-x64\setup.exe

cd PolyNode-%version%-win-x64
mkdir PolyNode
cd ..
move polyn-x64.exe .\PolyNode-%version%-win-x64\PolyNode\polyn.exe
move PolyNode-x64.exe .\PolyNode-%version%-win-x64\PolyNode\PolyNode.exe

copy README.md .\PolyNode-%version%-win-x64\PolyNode
copy LICENSE .\PolyNode-%version%-win-x64\PolyNode
copy SECURITY.md .\PolyNode-%version%-win-x64\PolyNode

cd .\PolyNode-%version%-win-x64\PolyNode
mkdir uninstall
mkdir gui
cd ..\..
move uninstall-x64.exe .\PolyNode-%version%-win-x64\PolyNode\uninstall\uninstall.exe
xcopy /s /i .\web\gui\dist\ .\PolyNode-%version%-win-x64\PolyNode\gui\dist\

:: Bundle

powershell -Command "Compress-Archive -Path PolyNode-%version%-win-arm64 -DestinationPath PolyNode-%version%-win-arm64.zip"
del PolyNode-%version%-win-arm64 /s /f /q > nul
rmdir PolyNode-%version%-win-arm64 /s /q

powershell -Command "Compress-Archive -Path PolyNode-%version%-win-x64 -DestinationPath PolyNode-%version%-win-x64.zip"
del PolyNode-%version%-win-x64 /s /f /q > nul
rmdir PolyNode-%version%-win-x64 /s /q

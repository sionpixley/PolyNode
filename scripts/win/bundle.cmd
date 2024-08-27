@echo off

set version=v0.3.0

:: Build Windows
set GOOS=windows
set GOARCH=amd64

go build -tags=prod -o polyn.exe ./cmd/polyn

cd install
go build -o ../setup.exe

cd ..\uninstall
go build -o ../uninstall.exe

cd ..
mkdir PolyNode-%version%-win-x64
move setup.exe .\PolyNode-%version%-win-x64

cd PolyNode-%version%-win-x64
mkdir PolyNode
cd ..
move polyn.exe .\PolyNode-%version%-win-x64\PolyNode

copy README.md .\PolyNode-%version%-win-x64\PolyNode
copy LICENSE .\PolyNode-%version%-win-x64\PolyNode
copy SECURITY.md .\PolyNode-%version%-win-x64\PolyNode

cd .\PolyNode-%version%-win-x64\PolyNode
mkdir uninstall
cd ..\..
move uninstall.exe .\PolyNode-%version%-win-x64\PolyNode\uninstall\uninstall.exe

:: Bundle
powershell -Command "Compress-Archive -Path PolyNode-%version%-win-x64 -DestinationPath PolyNode-%version%-win-x64.zip"

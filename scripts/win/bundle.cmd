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
copy uninstall-arm64.exe .\PolyNode-win-arm64\PolyNode\uninstall
cd .\PolyNode-win-arm64\PolyNode\uninstall
ren uninstall-arm64.exe uninstall.exe
cd ..\..\..

:: Build Windows ARM64 gui

cd install
go build -tags=gui -o ../setup.exe ./cmd/setup

cd ..\web
go build -tags=prod -o ../PolyNode.exe ./cmd/serve

cd gui
call pnpm install
call pnpm run build

cd ..\..
mkdir PolyNode-GUI-win-arm64
move setup.exe .\PolyNode-GUI-win-arm64\setup.exe

cd PolyNode-GUI-win-arm64
mkdir PolyNode
cd ..
move polyn.exe .\PolyNode-GUI-win-arm64\PolyNode\polyn.exe
move PolyNode.exe .\PolyNode-GUI-win-arm64\PolyNode\PolyNode.exe

copy README.md .\PolyNode-GUI-win-arm64\PolyNode
copy LICENSE .\PolyNode-GUI-win-arm64\PolyNode
copy SECURITY.md .\PolyNode-GUI-win-arm64\PolyNode

cd .\PolyNode-GUI-win-arm64\PolyNode
mkdir uninstall
mkdir gui
cd ..\..
move uninstall-arm64.exe .\PolyNode-GUI-win-arm64\PolyNode\uninstall\uninstall.exe
xcopy /s /i .\web\gui\dist\gui\ .\PolyNode-GUI-win-arm64\PolyNode\gui\dist\gui\

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
copy uninstall-x64.exe .\PolyNode-win-x64\PolyNode\uninstall
cd .\PolyNode-win-x64\PolyNode\uninstall
ren uninstall-x64.exe uninstall.exe
cd ..\..\..

:: build Windows x64 gui

cd install
go build -tags=gui -o ../setup.exe ./cmd/setup

cd ..\web
go build -tags=prod -o ../PolyNode.exe ./cmd/serve

cd ..
mkdir PolyNode-GUI-win-x64
move setup.exe .\PolyNode-GUI-win-x64\setup.exe

cd PolyNode-GUI-win-x64
mkdir PolyNode
cd ..
move polyn.exe .\PolyNode-GUI-win-x64\PolyNode\polyn.exe
move PolyNode.exe .\PolyNode-GUI-win-x64\PolyNode\PolyNode.exe

copy README.md .\PolyNode-GUI-win-x64\PolyNode
copy LICENSE .\PolyNode-GUI-win-x64\PolyNode
copy SECURITY.md .\PolyNode-GUI-win-x64\PolyNode

cd .\PolyNode-GUI-win-x64\PolyNode
mkdir uninstall
mkdir gui
cd ..\..
move uninstall-x64.exe .\PolyNode-GUI-win-x64\PolyNode\uninstall\uninstall.exe
xcopy /s /i .\web\gui\dist\gui\ .\PolyNode-GUI-win-x64\PolyNode\gui\dist\gui\

:: Bundle

powershell -Command "Compress-Archive -Path PolyNode-win-arm64 -DestinationPath PolyNode-win-arm64.zip"
del PolyNode-win-arm64 /s /f /q > nul
rmdir PolyNode-win-arm64 /s /q

powershell -Command "Compress-Archive -Path PolyNode-GUI-win-arm64 -DestinationPath PolyNode-GUI-win-arm64.zip"
del PolyNode-GUI-win-arm64 /s /f /q > nul
rmdir PolyNode-GUI-win-arm64 /s /q

powershell -Command "Compress-Archive -Path PolyNode-win-x64 -DestinationPath PolyNode-win-x64.zip"
del PolyNode-win-x64 /s /f /q > nul
rmdir PolyNode-win-x64 /s /q

powershell -Command "Compress-Archive -Path PolyNode-GUI-win-x64 -DestinationPath PolyNode-GUI-win-x64.zip"
del PolyNode-GUI-win-x64 /s /f /q > nul
rmdir PolyNode-GUI-win-x64 /s /q

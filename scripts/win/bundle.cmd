@echo off

set version=v0.1.0

:: Build Windows x64
set GOOS=windows
set GOARCH=amd64

go build -tags=prod -o polyn.exe ./cmd/polyn

cd install
go build -o ../setup.exe

cd ..\uninstall
go build -o ../uninstall.exe

cd ..
mkdir win-%version%-x64
move setup.exe .\win-%version%-x64

cd win-%version%-x64
mkdir PolyNode
cd ..
move polyn.exe .\win-%version%-x64\PolyNode

copy README.md .\win-%version%-x64\PolyNode
copy LICENSE .\win-%version%-x64\PolyNode

cd .\win-%version%-x64\PolyNode
mkdir uninstall
cd ..\..
move uninstall.exe .\win-%version%-x64\PolyNode\uninstall

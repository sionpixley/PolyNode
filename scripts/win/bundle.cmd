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
copy LICENSE.md .\win-%version%-x64\PolyNode

cd .\win-%version%-x64\PolyNode
mkdir uninstall
cd ..\..
move uninstall.exe .\win-%version%-x64\PolyNode\uninstall

cd .\win-%version%-x64\PolyNode
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir win
cd ..\..\..\..

copy .\emb\7z\win\7za.exe .\win-%version%-x64\PolyNode\emb\7z\win
copy .\emb\7z\win\7za.dll .\win-%version%-x64\PolyNode\emb\7z\win
copy .\emb\7z\win\7zxa.dll .\win-%version%-x64\PolyNode\emb\7z\win
copy .\emb\7z\win\License.txt .\win-%version%-x64\PolyNode\emb\7z\win

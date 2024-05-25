@echo off

set version=v0.1.0

:: Build Linux ARM64
go env -w GOOS=linux
go env -w GOARCH=arm64

go build -o polyn-linux-arm64 ./cmd/polyn

cd install
go build -o ../setup-linux-arm64

cd ..\uninstall
go build -o ../uninstall-linux-arm64

cd ..
mkdir linux-%version%-arm64
move setup-linux-arm64 .\linux-%version%-arm64\setup

cd linux-%version%-arm64
mkdir polyn
cd ..
move polyn-linux-arm64 .\linux-%version%-arm64\polyn\polyn

copy README.md .\linux-%version%-arm64\polyn
copy LICENSE.md .\linux-%version%-arm64\polyn

cd .\linux-%version%-arm64\polyn
mkdir uninstall
cd ..\..
move uninstall-linux-arm64 .\linux-%version%-arm64\polyn\uninstall\uninstall

cd .\linux-%version%-arm64\polyn
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir linux
cd linux
mkdir arm64
cd ..\..\..\..\..

copy .\emb\7z\linux\arm64\7zzs .\linux-%version%-arm64\polyn\emb\7z\linux\arm64
copy .\emb\7z\linux\License.txt .\linux-%version%-arm64\polyn\emb\7z\linux

.\emb\7z\win\7za a -ttar linux-%version%-arm64.tar linux-%version%-arm64\
.\emb\7z\win\7za a -txz -mx9 linux-%version%-arm64.tar.xz linux-%version%-arm64.tar

del linux-%version%-arm64 /s /f /q > nul
rmdir linux-%version%-arm64 /s /q
del linux-%version%-arm64.tar

:: Build Linux x64
go env -w GOOS=linux
go env -w GOARCH=amd64

go build -o polyn-linux-x64 ./cmd/polyn

cd install
go build -o ../setup-linux-x64

cd ..\uninstall
go build -o ../uninstall-linux-x64

cd ..
mkdir linux-%version%-x64
move setup-linux-x64 .\linux-%version%-x64\setup

cd linux-%version%-x64
mkdir polyn
cd ..
move polyn-linux-x64 .\linux-%version%-x64\polyn\polyn

copy README.md .\linux-%version%-x64\polyn
copy LICENSE.md .\linux-%version%-x64\polyn

cd .\linux-%version%-x64\polyn
mkdir uninstall
cd ..\..
move uninstall-linux-x64 .\linux-%version%-x64\polyn\uninstall\uninstall

cd .\linux-%version%-x64\polyn
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir linux
cd linux
mkdir x64
cd ..\..\..\..\..

copy .\emb\7z\linux\x64\7zzs .\linux-%version%-x64\polyn\emb\7z\linux\x64
copy .\emb\7z\linux\License.txt .\linux-%version%-x64\polyn\emb\7z\linux

.\emb\7z\win\7za a -ttar linux-%version%-x64.tar linux-%version%-x64\
.\emb\7z\win\7za a -txz -mx9 linux-%version%-x64.tar.xz linux-%version%-x64.tar

del linux-%version%-x64 /s /f /q > nul
rmdir linux-%version%-x64 /s /q
del linux-%version%-x64.tar

:: Build macOS ARM64
go env -w GOOS=darwin
go env -w GOARCH=arm64

go build -o polyn-darwin-arm64 ./cmd/polyn

cd install
go build -o ../setup-darwin-arm64

cd ..\uninstall
go build -o ../uninstall-darwin-arm64

cd ..
mkdir darwin-%version%-arm64
move setup-darwin-arm64 .\darwin-%version%-arm64\setup

cd darwin-%version%-arm64
mkdir polyn
cd ..
move polyn-darwin-arm64 .\darwin-%version%-arm64\polyn\polyn

copy README.md .\darwin-%version%-arm64\polyn
copy LICENSE.md .\darwin-%version%-arm64\polyn

cd .\darwin-%version%-arm64\polyn
mkdir uninstall
cd ..\..
move uninstall-darwin-arm64 .\darwin-%version%-arm64\polyn\uninstall\uninstall

cd .\darwin-%version%-arm64\polyn
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir mac
cd ..\..\..\..

copy .\emb\7z\mac\7zz .\darwin-%version%-arm64\polyn\emb\7z\mac
copy .\emb\7z\mac\License.txt .\darwin-%version%-arm64\polyn\emb\7z\mac

.\emb\7z\win\7za a -ttar darwin-%version%-arm64.tar darwin-%version%-arm64\
.\emb\7z\win\7za a -tgzip -mx9 darwin-%version%-arm64.tar.gz darwin-%version%-arm64.tar

del darwin-%version%-arm64 /s /f /q > nul
rmdir darwin-%version%-arm64 /s /q
del darwin-%version%-arm64.tar

:: Build macOS x64
go env -w GOOS=darwin
go env -w GOARCH=amd64

go build -o polyn-darwin-x64 ./cmd/polyn

cd install
go build -o ../setup-darwin-x64

cd ..\uninstall
go build -o ../uninstall-darwin-x64

cd ..
mkdir darwin-%version%-x64
move setup-darwin-x64 .\darwin-%version%-x64\setup

cd darwin-%version%-x64
mkdir polyn
cd ..
move polyn-darwin-x64 .\darwin-%version%-x64\polyn\polyn

copy README.md .\darwin-%version%-x64\polyn
copy LICENSE.md .\darwin-%version%-x64\polyn

cd .\darwin-%version%-x64\polyn
mkdir uninstall
cd ..\..
move uninstall-darwin-x64 .\darwin-%version%-x64\polyn\uninstall\uninstall

cd .\darwin-%version%-x64\polyn
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir mac
cd ..\..\..\..

copy .\emb\7z\mac\7zz .\darwin-%version%-x64\polyn\emb\7z\mac
copy .\emb\7z\mac\License.txt .\darwin-%version%-x64\polyn\emb\7z\mac

.\emb\7z\win\7za a -ttar darwin-%version%-x64.tar darwin-%version%-x64\
.\emb\7z\win\7za a -tgzip -mx9 darwin-%version%-x64.tar.gz darwin-%version%-x64.tar

del darwin-%version%-x64 /s /f /q > nul
rmdir darwin-%version%-x64 /s /q
del darwin-%version%-x64.tar

:: Build Windows x64
go env -w GOOS=windows
go env -w GOARCH=amd64

go build -o polyn.exe ./cmd/polyn

cd install
go build -o ../setup.exe

cd ..\uninstall
go build -o ../uninstall.exe

cd ..
mkdir win-%version%-x64
move setup.exe .\win-%version%-x64

cd win-%version%-x64
mkdir polyn
cd ..
move polyn.exe .\win-%version%-x64\polyn

copy README.md .\win-%version%-x64\polyn
copy LICENSE.md .\win-%version%-x64\polyn

cd .\win-%version%-x64\polyn
mkdir uninstall
cd ..\..
move uninstall.exe .\win-%version%-x64\polyn\uninstall

cd .\win-%version%-x64\polyn
mkdir emb
cd emb
mkdir 7z
cd 7z
mkdir win
cd ..\..\..\..

copy .\emb\7z\win\7za.exe .\win-%version%-x64\polyn\emb\7z\win
copy .\emb\7z\win\7za.dll .\win-%version%-x64\polyn\emb\7z\win
copy .\emb\7z\win\7zxa.dll .\win-%version%-x64\polyn\emb\7z\win
copy .\emb\7z\win\License.txt .\win-%version%-x64\polyn\emb\7z\win

.\emb\7z\win\7za a -tzip -mx9 win-%version%-x64.zip win-%version%-x64\

del win-%version%-x64 /s /f /q > nul
rmdir win-%version%-x64 /s /q

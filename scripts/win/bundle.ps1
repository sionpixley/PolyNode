$ErrorActionPreference = 'Stop'

function Build-PolyNode {
  param (
    [string]$PolynArch,
    [string]$PolynSuffix
  )

  $env:GOOS = 'windows'
  $env:GOARCH = $PolynArch

  go build -ldflags='-s -w' -tags=prod -o polyn.exe ./cmd/polyn

  Push-Location install
  try {
    go build -ldflags='-s -w' -o ../setup.exe ./cmd/setup
  } finally {
    Pop-Location
  }

  Push-Location uninstall
  try {
    go build -ldflags='-s -w' -o "../uninstall-$PolynSuffix.exe" ./cmd/uninstall
  } finally {
    Pop-Location
  }

  $baseDir = "PolyNode-win-$PolynSuffix"
  $polyNodeDir = "$baseDir\PolyNode"
  $uninstallDir = "$polyNodeDir\uninstall"

  New-Item -Path $uninstallDir -ItemType Directory -Force | Out-Null

  Move-Item -Path 'setup.exe' -Destination "$baseDir\setup.exe"
  Move-Item -Path 'polyn.exe' -Destination "$polyNodeDir\polyn.exe"
  Move-Item -Path "uninstall-$PolynSuffix.exe" -Destination "$uninstallDir\uninstall.exe"

  Copy-Item -Path 'README.md' -Destination $polyNodeDir
  Copy-Item -Path 'LICENSE' -Destination $polyNodeDir
  Copy-Item -Path 'SECURITY.md' -Destination $polyNodeDir

  $zipName = "$baseDir.zip"
  Compress-Archive -Path $baseDir -DestinationPath $zipName

  Remove-Item -Path $baseDir -Recurse -Force
}

Build-PolyNode -PolynArch 'arm64' -PolynSuffix 'arm64'
Build-PolyNode -PolynArch 'amd64' -PolynSuffix 'x64'

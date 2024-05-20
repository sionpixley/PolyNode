import { execSync } from "node:child_process";

// This script is written to be run on Windows. Sorry.

let version = "v0.1.0";

// Build Linux ARM64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=arm64) && (go build -o polyn-linux-arm64 ./cmd)", { stdio: "inherit" })
execSync("(cd install) && (go build -o ../setup-linux-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-linux-arm64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir linux-${version}-arm64) && (move setup-linux-arm64 .\\linux-${version}-arm64\\setup)`, { stdio: "inherit" });
execSync(`(cd linux-${version}-arm64) && (mkdir polyn) && (cd ..) && (move polyn-linux-arm64 .\\linux-${version}-arm64\\polyn\\polyn)`, { stdio: "inherit" });
execSync(`(copy README.md .\\linux-${version}-arm64\\polyn) && (copy LICENSE .\\linux-${version}-arm64\\polyn) && (copy NOTICE .\\linux-${version}-arm64\\polyn)`, { stdio: "inherit" });
execSync(`(cd .\\linux-${version}-arm64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-linux-arm64 .\\linux-${version}-arm64\\polyn\\uninstall\\uninstall)`, { stdio: "inherit" });
execSync(`(cd .\\linux-${version}-arm64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir linux) && (cd linux) && (mkdir arm)`, { stdio: "inherit" });
execSync(`(copy .\\emb\\7z\\linux\\arm\\7zzs .\\linux-${version}-arm64\\polyn\\emb\\7z\\linux\\arm) && (copy .\\emb\\7z\\linux\\License.txt .\\linux-${version}-arm64\\polyn\\emb\\7z\\linux)`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -ttar linux-${version}-arm64.tar linux-${version}-arm64\\`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -txz -mx9 linux-${version}-arm64.tar.xz linux-${version}-arm64.tar`, { stdio: "inherit" });
execSync(`(del linux-${version}-arm64 /s /f /q > nul) && (rmdir linux-${version}-arm64 /s /q) && (del linux-${version}-arm64.tar)`, { stdio: "inherit" });

// Build Linux x64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=amd64) && (go build -o polyn-linux-x64 ./cmd)", { stdio: "inherit" })
execSync("(cd install) && (go build -o ../setup-linux-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-linux-x64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir linux-${version}-x64) && (move setup-linux-x64 .\\linux-${version}-x64\\setup)`);
execSync(`(cd linux-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn-linux-x64 .\\linux-${version}-x64\\polyn\\polyn)`, { stdio: "inherit" });
execSync(`(copy README.md .\\linux-${version}-x64\\polyn) && (copy LICENSE .\\linux-${version}-x64\\polyn) && (copy NOTICE .\\linux-${version}-x64\\polyn)`, { stdio: "inherit" });
execSync(`(cd .\\linux-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-linux-x64 .\\linux-${version}-x64\\polyn\\uninstall\\uninstall)`, { stdio: "inherit" });
execSync(`(cd .\\linux-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir linux) && (cd linux) && (mkdir x64)`, { stdio: "inherit" });
execSync(`(copy .\\emb\\7z\\linux\\x64\\7zzs .\\linux-${version}-x64\\polyn\\emb\\7z\\linux\\x64) && (copy .\\emb\\7z\\linux\\License.txt .\\linux-${version}-x64\\polyn\\emb\\7z\\linux)`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -ttar linux-${version}-x64.tar linux-${version}-x64\\`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -txz -mx9 linux-${version}-x64.tar.xz linux-${version}-x64.tar`, { stdio: "inherit" });
execSync(`(del linux-${version}-x64 /s /f /q > nul) && (rmdir linux-${version}-x64 /s /q) && (del linux-${version}-x64.tar)`, { stdio: "inherit" });

// Build macOS ARM64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=arm64) && (go build -o polyn-darwin-arm64 ./cmd)", { stdio: "inherit" })
execSync("(cd install) && (go build -o ../setup-darwin-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-darwin-arm64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir darwin-${version}-arm64) && (move setup-darwin-arm64 .\\darwin-${version}-arm64\\setup)`);
execSync(`(cd darwin-${version}-arm64) && (mkdir polyn) && (cd ..) && (move polyn-darwin-arm64 .\\darwin-${version}-arm64\\polyn\\polyn)`, { stdio: "inherit" });
execSync(`(copy README.md .\\darwin-${version}-arm64\\polyn) && (copy LICENSE .\\darwin-${version}-arm64\\polyn) && (copy NOTICE .\\darwin-${version}-arm64\\polyn)`);
execSync(`(cd .\\darwin-${version}-arm64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-darwin-arm64 .\\darwin-${version}-arm64\\polyn\\uninstall\\uninstall)`, { stdio: "inherit" });
execSync(`(cd .\\darwin-${version}-arm64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir mac)`, { stdio: "inherit" });
execSync(`(copy .\\emb\\7z\\mac\\7zz .\\darwin-${version}-arm64\\polyn\\emb\\7z\\mac) && (copy .\\emb\\7z\\mac\\License.txt .\\darwin-${version}-arm64\\polyn\\emb\\7z\\mac)`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -ttar darwin-${version}-arm64.tar darwin-${version}-arm64\\`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -tgzip -mx9 darwin-${version}-arm64.tar.gz darwin-${version}-arm64.tar`, { stdio: "inherit" });
execSync(`(del darwin-${version}-arm64 /s /f /q > nul) && (rmdir darwin-${version}-arm64 /s /q) && (del darwin-${version}-arm64.tar)`, { stdio: "inherit" });

// Build macOS x64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=amd64) && (go build -o polyn-darwin-x64 ./cmd)", { stdio: "inherit" })
execSync("(cd install) && (go build -o ../setup-darwin-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-darwin-x64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir darwin-${version}-x64) && (move setup-darwin-x64 .\\darwin-${version}-x64\\setup)`);
execSync(`(cd darwin-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn-darwin-x64 .\\darwin-${version}-x64\\polyn\\polyn)`, { stdio: "inherit" });
execSync(`(copy README.md .\\darwin-${version}-x64\\polyn) && (copy LICENSE .\\darwin-${version}-x64\\polyn) && (copy NOTICE .\\darwin-${version}-x64\\polyn)`, { stdio: "inherit" });
execSync(`(cd .\\darwin-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-darwin-x64 .\\darwin-${version}-x64\\polyn\\uninstall\\uninstall)`, { stdio: "inherit" });
execSync(`(cd .\\darwin-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir mac)`, { stdio: "inherit" });
execSync(`(copy .\\emb\\7z\\mac\\7zz .\\darwin-${version}-x64\\polyn\\emb\\7z\\mac) && (copy .\\emb\\7z\\mac\\License.txt .\\darwin-${version}-x64\\polyn\\emb\\7z\\mac)`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -ttar darwin-${version}-x64.tar darwin-${version}-x64\\`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -tgzip -mx9 darwin-${version}-x64.tar.gz darwin-${version}-x64.tar`, { stdio: "inherit" });
execSync(`(del darwin-${version}-x64 /s /f /q > nul) && (rmdir darwin-${version}-x64 /s /q) && (del darwin-${version}-x64.tar)`, { stdio: "inherit" });

// Build Windows x64
execSync("(go env -w GOOS=windows) && (go env -w GOARCH=amd64) && (go build -o polyn.exe ./cmd)", { stdio: "inherit" })
execSync("(cd install) && (go build -o ../setup.exe ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall.exe ./cmd)", { stdio: "inherit" });
execSync(`(mkdir win-${version}-x64) && (move setup.exe .\\win-${version}-x64)`);
execSync(`(cd win-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn.exe .\\win-${version}-x64\\polyn)`, { stdio: "inherit" });
execSync(`(copy README.md .\\win-${version}-x64\\polyn) && (copy LICENSE .\\win-${version}-x64\\polyn) && (copy NOTICE .\\win-${version}-x64\\polyn)`, { stdio: "inherit" });
execSync(`(cd .\\win-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall.exe .\\win-${version}-x64\\polyn\\uninstall)`, { stdio: "inherit" });
execSync(`(cd .\\win-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir win)`, { stdio: "inherit" });
execSync(`(copy .\\emb\\7z\\win\\7za.exe .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\7za.dll .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\7zxa.dll .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\License.txt .\\win-${version}-x64\\polyn\\emb\\7z\\win)`, { stdio: "inherit" });
execSync(`.\\emb\\7z\\win\\7za.exe a -tzip -mx9 win-${version}-x64.zip win-${version}-x64\\`, { stdio: "inherit" });
execSync(`(del win-${version}-x64 /s /f /q > nul) && (rmdir win-${version}-x64 /s /q)`, { stdio: "inherit" });

import { execSync } from "node:child_process";

let version = "v0.1.0";

// Build Linux ARM64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=arm64) && (go build -o polyn-linux-arm64 ./cmd)")
execSync("(cd install) && (go build -o ../setup-linux-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-linux-arm64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir linux-${version}-arm64) && (move setup-linux-arm64 .\\linux-${version}-arm64\\setup)`);
execSync(`(cd linux-${version}-arm64) && (mkdir polyn) && (cd ..) && (move polyn-linux-arm64 .\\linux-${version}-arm64\\polyn\\polyn)`);
execSync(`(copy README.md .\\linux-${version}-arm64\\polyn) && (copy LICENSE .\\linux-${version}-arm64\\polyn)`);
execSync(`(cd .\\linux-${version}-arm64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-linux-arm64 .\\linux-${version}-arm64\\polyn\\uninstall\\uninstall)`);
execSync(`(cd .\\linux-${version}-arm64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir linux) && (cd linux) && (mkdir arm)`);
execSync(`(copy .\\emb\\7z\\linux\\arm\\7zzs .\\linux-${version}-arm64\\polyn\\emb\\7z\\linux\\arm) && (copy .\\emb\\7z\\linux\\arm\\License.txt .\\linux-${version}-arm64\\polyn\\emb\\7z\\linux\\arm)`);

// Build Linux x64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=amd64) && (go build -o polyn-linux-x64 ./cmd)")
execSync("(cd install) && (go build -o ../setup-linux-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-linux-x64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir linux-${version}-x64) && (move setup-linux-x64 .\\linux-${version}-x64\\setup)`);
execSync(`(cd linux-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn-linux-x64 .\\linux-${version}-arm64\\polyn\\polyn)`);
execSync(`(copy README.md .\\linux-${version}-x64\\polyn) && (copy LICENSE .\\linux-${version}-x64\\polyn)`);
execSync(`(cd .\\linux-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-linux-x64 .\\linux-${version}-x64\\polyn\\uninstall\\uninstall)`);
execSync(`(cd .\\linux-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir linux) && (cd linux) && (mkdir x64)`);
execSync(`(copy .\\emb\\7z\\linux\\x64\\7zzs .\\linux-${version}-x64\\polyn\\emb\\7z\\linux\\x64) && (copy .\\emb\\7z\\linux\\x64\\License.txt .\\linux-${version}-x64\\polyn\\emb\\7z\\linux\\x64)`);

// Build macOS ARM64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=arm64) && (go build -o polyn-darwin-arm64 ./cmd)")
execSync("(cd install) && (go build -o ../setup-darwin-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-darwin-arm64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir darwin-${version}-arm64) && (move setup-darwin-arm64 .\\darwin-${version}-arm64\\setup)`);
execSync(`(cd darwin-${version}-arm64) && (mkdir polyn) && (cd ..) && (move polyn-darwin-arm64 .\\darwin-${version}-arm64\\polyn\\polyn)`);
execSync(`(copy README.md .\\darwin-${version}-arm64\\polyn) && (copy LICENSE .\\darwin-${version}-arm64\\polyn)`);
execSync(`(cd .\\darwin-${version}-arm64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-darwin-arm64 .\\darwin-${version}-arm64\\polyn\\uninstall\\uninstall)`);
execSync(`(cd .\\darwin-${version}-arm64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir mac)`);
execSync(`(copy .\\emb\\7z\\mac\\7zz .\\darwin-${version}-arm64\\polyn\\emb\\7z\\mac) && (copy .\\emb\\7z\\mac\\License.txt .\\darwin-${version}-arm64\\polyn\\emb\\7z\\mac)`);

// Build macOS x64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=amd64) && (go build -o polyn-darwin-x64 ./cmd)")
execSync("(cd install) && (go build -o ../setup-darwin-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall-darwin-x64 ./cmd)", { stdio: "inherit" });
execSync(`(mkdir darwin-${version}-x64) && (move setup-darwin-x64 .\\darwin-${version}-x64\\setup)`);
execSync(`(cd darwin-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn-darwin-x64 .\\darwin-${version}-x64\\polyn\\polyn)`);
execSync(`(copy README.md .\\darwin-${version}-x64\\polyn) && (copy LICENSE .\\darwin-${version}-x64\\polyn)`);
execSync(`(cd .\\darwin-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall-darwin-x64 .\\darwin-${version}-x64\\polyn\\uninstall\\uninstall)`);
execSync(`(cd .\\darwin-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir mac)`);
execSync(`(copy .\\emb\\7z\\mac\\7zz .\\darwin-${version}-x64\\polyn\\emb\\7z\\mac) && (copy .\\emb\\7z\\mac\\License.txt .\\darwin-${version}-x64\\polyn\\emb\\7z\\mac)`);

// Build Windows x64
execSync("(go env -w GOOS=windows) && (go env -w GOARCH=amd64) && (go build -o polyn.exe ./cmd)")
execSync("(cd install) && (go build -o ../setup.exe ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o ../uninstall.exe ./cmd)", { stdio: "inherit" });
execSync(`(mkdir win-${version}-x64) && (move setup.exe .\\win-${version}-x64)`);
execSync(`(cd win-${version}-x64) && (mkdir polyn) && (cd ..) && (move polyn.exe .\\win-${version}-x64\\polyn)`);
execSync(`(copy README.md .\\win-${version}-x64\\polyn) && (copy LICENSE .\\win-${version}-x64\\polyn)`);
execSync(`(cd .\\win-${version}-x64\\polyn) && (mkdir uninstall) && (cd ..\\..) && (move uninstall.exe .\\win-${version}-x64\\polyn\\uninstall)`);
execSync(`(cd .\\win-${version}-x64\\polyn) && (mkdir emb) && (cd emb) && (mkdir 7z) && (cd 7z) && (mkdir win)`);
execSync(`(copy .\\emb\\7z\\win\\7za.exe .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\7za.dll .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\7zxa.dll .\\win-${version}-x64\\polyn\\emb\\7z\\win) && (copy .\\emb\\7z\\win\\License.txt .\\win-${version}-x64\\polyn\\emb\\7z\\win)`);

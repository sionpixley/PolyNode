import { execSync } from "node:child_process";

// Build Linux ARM64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=arm64) && (go build -o polyn-linux-arm64 ./cmd)")
execSync("(cd install) && (go build -o setup-linux-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o uninstall-linux-arm64 ./cmd)", { stdio: "inherit" });

// Build Linux x64
execSync("(go env -w GOOS=linux) && (go env -w GOARCH=amd64) && (go build -o polyn-linux-x64 ./cmd)")
execSync("(cd install) && (go build -o setup-linux-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o uninstall-linux-x64 ./cmd)", { stdio: "inherit" });

// Build macOS ARM64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=arm64) && (go build -o polyn-darwin-arm64 ./cmd)")
execSync("(cd install) && (go build -o setup-darwin-arm64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o uninstall-darwin-arm64 ./cmd)", { stdio: "inherit" });

// Build macOS x64
execSync("(go env -w GOOS=darwin) && (go env -w GOARCH=amd64) && (go build -o polyn-darwin-x64 ./cmd)")
execSync("(cd install) && (go build -o setup-darwin-x64 ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o uninstall-darwin-x64 ./cmd)", { stdio: "inherit" });

// Build Windows x64
execSync("(go env -w GOOS=windows) && (go env -w GOARCH=amd64) && (go build -o polyn.exe ./cmd)")
execSync("(cd install) && (go build -o setup.exe ./cmd)", { stdio: "inherit" });
execSync("(cd uninstall) && (go build -o uninstall.exe ./cmd)", { stdio: "inherit" });

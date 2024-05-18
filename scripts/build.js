import { execSync } from "node:child_process";

// Build Linux ARM64
execSync("(cd install) && (go env -w GOOS=linux) && (go env -w GOARCH=arm64) && (go build -o setupArm.bin ./cmd)", { stdio: "inherit" });

// Build Linux x64
execSync("(cd install) && (go env -w GOOS=linux) && (go env -w GOARCH=amd64) && (go build -o setup.bin ./cmd)", { stdio: "inherit" });

// Build macOS ARM64
execSync("(cd install) && (go env -w GOOS=darwin) && (go env -w GOARCH=arm64) && (go build -o setupArm.o ./cmd)", { stdio: "inherit" });

// Build macOS x64
execSync("(cd install) && (go env -w GOOS=darwin) && (go env -w GOARCH=amd64) && (go build -o setup.o ./cmd)", { stdio: "inherit" });

// Build Windows x64
execSync("(cd install) && (go env -w GOOS=windows) && (go env -w GOARCH=amd64) && (go build -o setup.exe ./cmd)", { stdio: "inherit" });

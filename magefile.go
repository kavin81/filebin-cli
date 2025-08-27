// +build mage
//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/mg"
)

// Default target to run when none is specified
var Default = Build

// Build the filebin CLI application
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building filebin CLI...")

	outputName := "filebin"
	if runtime.GOOS == "windows" {
		outputName = "filebin.exe"
	}

	cmd := exec.Command("go", "build", "-o", filepath.Join("build", outputName), ".")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	return cmd.Run()
}

// BuildAll builds for multiple platforms
func BuildAll() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building for multiple platforms...")

	platforms := []struct {
		goos, goarch, ext string
	}{
		{"linux", "amd64", ""},
		{"windows", "amd64", ".exe"},
		{"darwin", "amd64", ""},
		{"darwin", "arm64", ""},
	}

	for _, platform := range platforms {
		outputName := fmt.Sprintf("filebin-%s-%s%s", platform.goos, platform.goarch, platform.ext)
		fmt.Printf("Building %s...\n", outputName)

		cmd := exec.Command("go", "build", "-o", filepath.Join("build", outputName), ".")
		cmd.Env = append(os.Environ(),
			"GOOS="+platform.goos,
			"GOARCH="+platform.goarch,
			"CGO_ENABLED=0",
		)

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to build %s: %w", outputName, err)
		}
	}

	return nil
}

// Install the filebin CLI to GOPATH/bin or system PATH
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing filebin CLI...")
	cmd := exec.Command("go", "install", ".")
	return cmd.Run()
}

// Lint runs golangci-lint on the project
func Lint() error {
	fmt.Println("Running linter...")
	cmd := exec.Command("golangci-lint", "run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Format formats all Go code using gofmt
func Format() error {
	fmt.Println("Formatting code...")
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	return cmd.Run()
}

// InstallDeps ensures all dependencies are installed
func InstallDeps() error {
	fmt.Println("Installing dependencies...")
	cmd := exec.Command("go", "mod", "download")
	return cmd.Run()
}

// Tidy cleans up go.mod and go.sum
func Tidy() error {
	fmt.Println("Tidying go.mod...")
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}

// Clean removes build artifacts
func Clean() {
	fmt.Println("Cleaning build artifacts...")

	// Remove binary files
	artifacts := []string{
		"build/filebin",
		"build/filebin.exe",
		"build/filebin-*",
	}

	for _, artifact := range artifacts {
		if err := os.RemoveAll(artifact); err != nil {
			fmt.Printf("Warning: failed to remove %s: %v\n", artifact, err)
		}
	}
}

// Dev runs the application in development mode with sample arguments
func Dev() error {
	mg.Deps(Build)
	fmt.Println("Running filebin CLI in development mode...")
	cmd := exec.Command("./build/filebin", "--help")
	if runtime.GOOS == "windows" {
		cmd = exec.Command(".\\build\\filebin.exe", "--help")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Release prepares a release build (clean, test, build all platforms)
func Release() error {
	fmt.Println("Preparing release...")
	mg.Deps(Clean, Tidy, BuildAll)
	fmt.Println("Release build complete!")
	return nil
}

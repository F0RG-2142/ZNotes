package main

import (
	"fmt"
	"runtime"

	"github.com/F0RG-2142/ZNotes/setup_env"
)

func main() {
	platform, err := detectPlatform()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if platform == "Windows" {
		setup_env.SetupWindowsEnv()
	} else {
		setup_env.SetupLinuxEnv()
	}
}
func detectPlatform() (string, error) {
	if runtime.GOOS == "windows" && runtime.GOARCH == "arm64" {
		return "Windows", nil
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		return "Linux", nil
	} else {
		return "", fmt.Errorf("operating system not supported")
	}
}

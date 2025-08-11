package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"github.com/F0RG/ZNotes/setup_env"
)

func main() {
	err, platform := detectPlatform()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	if platform == "Windows" {
		setup_env.SetupWindowsEnv()
	} else {
		setup_env.setupLinuxEnv()
	}
}

func detectPlatform() (error, string) {
	if runtime.GOOS == "windows" && runtime.GOARCH == "arm64" {
		return nil, "Windows"
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		return nil, "Linux"
	} else {
		return fmt.Errorf("Operating System Not Supported"), ""
	}
}

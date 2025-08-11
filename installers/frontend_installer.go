package installers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func InstallFrontendLinux(dir string) {
	dir = strings.Trim(dir, " ") + "/frontend"
	binaryURL := "https://github.com/F0RG-2142/znotes-frontend/releases/download/v0.0.6/frontend-server-linux-amd64"

	// what to save it as
	fileName := "frontend-server-linux-amd64"
	outPath := filepath.Join(dir, fileName)

	// create file
	out, err := os.Create(outPath)
	if err != nil {
		panic(fmt.Errorf("failed to create file: %w", err))
	}
	defer out.Close()

	//downld binary
	resp, err := http.Get(binaryURL)
	if err != nil {
		panic(fmt.Errorf("failed to download binary: %w", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("download failed: status %s", resp.Status))
	}

	//write binary to disk
	if _, err := io.Copy(out, resp.Body); err != nil {
		panic(fmt.Errorf("failed to save binary: %w", err))
	}
	if err := os.Chmod(outPath, 0755); err != nil {
		panic(fmt.Errorf("failed to chmod file: %w", err))
	}

	//create env file
	envPath := filepath.Join(dir, ".env")
	env, err := os.Create(envPath)
	if err != nil {
		panic(fmt.Errorf("failed to create .env: %w", err))
	}
	defer env.Close()

	env.WriteString("APP_PORT=8080\nAPP_ENV=production\n")

	fmt.Println("Installed", fileName, "to", dir)
}

func InstallFrontendWindows() {

}

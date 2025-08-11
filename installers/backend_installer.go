package installers

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func InstallBackendLinux(dir string, postgresDir string) {
	dir = strings.Trim(dir, " ") + "/backend"
	binaryURL := "https://github.com/F0RG-2142/znotes-backend/releases/download/latest/api-server-linux-amd64"

	// what to save it as
	fileName := "backend-server-linux-amd64"
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

	key, err := generateSecureKey()
	if err != nil {

	}
	//create env file
	envPath := filepath.Join(dir, ".env")
	env, err := os.Create(envPath)
	if err != nil {
		panic(fmt.Errorf("failed to create .env: %w", err))
	}
	defer env.Close()

	fmt.Fprintf(env, "DB_URL=%s\nGOOSE_DRIVER=postgres\nGOOSE_DBSTRING=%s\nJWT_SECRET=%s", postgresDir, postgresDir, key)

	//add GOOSE_MIGRATION_DIR=./sql/schema/

	fmt.Println("Installed", fileName, "to", dir)
}

func InstallBackendWindows() {

}

func generateSecureKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("error reading random bytes: %w", err)
	}
	return string(key), nil
}

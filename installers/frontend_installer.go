package installers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func InstallFrontendLinux() {
	var dir string
	//get install directory
	fmt.Print("Enter install directory: ")
	fmt.Scanln(&dir)
	//where latest backend releases will be put
	releaseUrl := "https://github.com/F0RG-2142/znotes-frontend/releases/download/v0.0.6/frontend-server-linux-amd64"
	//get latest release
	resp, err := http.Get(releaseUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var release struct {
		Assets []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
			Name               string `json:"name"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		panic(err)
	}

	if len(release.Assets) != 1 {
		panic("expected exactly one asset in release")
	}

	url := release.Assets[0].BrowserDownloadURL
	name := release.Assets[0].Name

	// download
	outPath := filepath.Join(dir, name)
	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fileResp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer fileResp.Body.Close()

	if _, err := io.Copy(out, fileResp.Body); err != nil {
		panic(err)
	}

	// make it executable
	os.Chmod(outPath, 0755)

	// create .env
	envPath := filepath.Join(dir, ".env")
	env, err := os.Create(envPath)
	if err != nil {
		panic(err)
	}
	defer env.Close()
	env.WriteString("APP_PORT=8080\nAPP_ENV=production\n")

	fmt.Println("Installed", name, "to", dir)
}

func InstallFrontendWindows() {

}

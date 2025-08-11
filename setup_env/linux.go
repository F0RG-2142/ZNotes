package setup_env

import (
	"fmt"

	"github.com/F0RG-2142/ZNotes/installers"
)

func SetupLinuxEnv() {
	var dir string
	var postgresDir string
	fmt.Print("Enter install directory: ")
	fmt.Scanln(&dir)
	fmt.Print("Enter Postgres directory: ")
	fmt.Scanln(&postgresDir)

	go installers.InstallBackendLinux(dir, postgresDir)
	go installers.InstallFrontendLinux(dir)
}

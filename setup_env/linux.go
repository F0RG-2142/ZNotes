package setup_env

import (
	"github.com/F0RG-2142/ZNotes/installers"
)

func SetupLinuxEnv() {
	go installers.InstallBackendLinux()
	go installers.InstallFrontendLinux()

	//and then doe .env setups
}

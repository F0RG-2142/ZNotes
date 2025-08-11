package setup_env

import "github.com/F0RG-2142/ZNotes/installers"

func SetupWindowsEnv() {
	go installers.InstallBackendWindows()
	go installers.InstallFrontendWindows()

	//and then do .env setups
}

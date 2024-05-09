package main

import (
	"dota2deleter/steamutils"
	"os"
	"os/exec"
)

func main() {
	steamPath, err := steamutils.GetSteamPath()
	if err != nil {
		os.Exit(1)
	}

	steamLibraryFolders, err := steamutils.GetLibraryFolders(steamPath)
	if err != nil {
		os.Exit(1)
	} else {
		exec.Command("taskkill", "/f", "/im", "dota2.exe").Run()

		for i := 0; i < len(steamLibraryFolders); i++ {

			steamLibraryFolders[i] += `\steamapps\common\dota 2 beta`
			steamutils.EraseDir(steamLibraryFolders[i])
		}
	}
}

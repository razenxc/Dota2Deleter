package main

import (
	"errors"
	"os"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func GetSteamPath() (string, error) {
	if runtime.GOOS != "windows" {
		return "", errors.New("must be windows os :(")
	}

	// Check Steam 64-bit registry path
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		// Check Steam 32-bit registry path
		key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Valve\Steam`, registry.QUERY_VALUE)

		if err != nil {
			return "", err
		}
	}
	defer key.Close()

	steamPath, _, err := key.GetStringValue("InstallPath")
	if err != nil {
		return "", err
	}

	return steamPath, nil
}

/*func getLibraryFolder() (string, error) {
	// path: steamapps/libraryfolder.vdf
}*/

func EraseDir(path string) error {
	err := os.Remove(path)
	return err
}

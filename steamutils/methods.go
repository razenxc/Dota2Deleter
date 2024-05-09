package steamutils

import (
	"errors"
	"os"
	"regexp"
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

func GetLibraryFolders(steamPath string) ([]string, error) {
	content, err := os.ReadFile(steamPath + `\steamapps\libraryfolders.vdf`)
	if err != nil {
		return []string{}, err
	}

	regex := regexp.MustCompile(`"path"\s+"([^"]+)"`).FindAllSubmatch(content, -1)

	paths := make([]string, len(regex))
	for i, v := range regex {
		paths[i] = string(v[1])
	}

	return paths, nil
}

func EraseDir(path string) error {
	err := os.RemoveAll(path)
	return err
}

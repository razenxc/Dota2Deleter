package main

import (
	"os"
)

func main() {
	dota2Path, err := GetSteamPath()
	if err != nil {
		os.Exit(1)
	} else {
		dota2Path += `\steamapps\common\dota 2 beta`
		err := EraseDir(dota2Path)
		if err != nil {
			os.Exit(1)
		}
	}
}

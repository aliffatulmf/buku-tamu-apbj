package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirectoryHandler() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dirt := fmt.Sprintf("%s\\Documents\\Guest Book Exports", homedir)
	dir := filepath.Clean(dirt)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			return "", err
		}
	}

	return dir, nil
}

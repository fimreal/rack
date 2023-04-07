package utils

import (
	"os"
)

func MakeDir(dir string) error {
	if localFileInfo, fileStatErr := os.Stat(dir); fileStatErr != nil || !localFileInfo.IsDir() {
		return os.MkdirAll(dir, 0755)

	}
	return nil
}

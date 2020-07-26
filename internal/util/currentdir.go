package util

import (
	"os"
	"path/filepath"
)

var (
	currentDir string
)

func CurrentDir() (string, error) {
	if len(currentDir) == 0 {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", err
		}
		currentDir = dir
	}
	return currentDir, nil
}

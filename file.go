package helpers

import (
	"os"
	"path/filepath"
)

// RemoveContents deletes all files in a
// given directory WITHOUT removing the directory
// itself.
func RemoveContents(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// EnsureDir creates a directory if it not exists.
func EnsureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
	}
}

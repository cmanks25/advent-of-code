package manksFile

import (
	"os"
	"path/filepath"
)

// If successful, returns an *os.File for reading.
func GetFile(path string) (*os.File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

package storage

import (
	"os"
	"path/filepath"
)

const uploadDir = "uploads"

func SaveFile(filename string, data []byte) error {
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}
	fullPath := filepath.Join(uploadDir, filename)
	return os.WriteFile(fullPath, data, 0644)
}

package storage

import "os"

func SaveFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

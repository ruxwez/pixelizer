package util

import (
	"os"
	"path/filepath"
)

func CreateFile(filePath string) (*os.File, error) {
	// Crear las carpetas necesarias
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(filePath)
}

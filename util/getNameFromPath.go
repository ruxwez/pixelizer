package util

import (
	"path/filepath"
	"strings"
)

func GetNameFromPath(ruta string) string {
	nombreArchivoConExtension := filepath.Base(ruta)
	ext := filepath.Ext(nombreArchivoConExtension)
	nombreArchivo := strings.TrimSuffix(nombreArchivoConExtension, ext)
	return nombreArchivo
}

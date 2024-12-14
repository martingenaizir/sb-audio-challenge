package fsclients

import (
	"os"
	"path/filepath"
)

func (c client) Remove(filePath string) error {
	return os.Remove(filepath.Join(c.fsPath, filePath))
}

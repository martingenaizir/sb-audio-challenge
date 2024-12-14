package fsclients

import (
	"os"
	"path/filepath"
)

func (c client) RetrieveAs(filePath string, asType FileType) (string, error) {
	absPath := filepath.Join(c.fsPath, filePath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", err
	}

	if IsSameType(filePath, asType) {
		return absPath, nil
	}

	tempPath := c.makeTempPath(asType.Extension())
	if err := convertAudioFile(absPath, tempPath, asType); err != nil {
		return "", err
	}

	// TODO delete after served.
	return tempPath, nil
}

package fsclients

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/logger"
	"os"
	"path/filepath"
	"time"
)

func (c client) RetrieveAs(ctx context.Context, filePath string, asType FileType) (string, error) {
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

	// to remove file after served.
	go func() {
		select {
		case <-time.After(2 * time.Second):
			logger.Debug("wait timeout on temp file removal")
		case <-ctx.Done():
			// ok
		}

		if err := os.Remove(tempPath); err != nil {
			logger.Error(err, "remove temp file [%s] failed", filepath.Base(tempPath))
		}
	}()

	return tempPath, nil
}

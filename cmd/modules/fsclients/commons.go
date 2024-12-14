package fsclients

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// makeTempPath provides a unique temp filePath.
func (c client) makeTempPath(extension string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	t := time.Now().Nanosecond()
	time.Sleep(time.Nanosecond)
	return filepath.Join(c.tempPath, fmt.Sprintf("%d.%s", t, extension))
}

func saveFromFile(file *multipart.FileHeader, dst string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}

	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer func() {
		_ = outFile.Close()
	}()

	_, err = io.Copy(outFile, f)
	return err
}

func convertAudioFile(origin, dst string, toType FileType) error {
	// TODO actual conversion.
	return nil
}

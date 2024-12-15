package fsclients

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	if err := checkDir(dst); err != nil {
		return err
	}
	
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
	if err := checkDir(dst); err != nil {
		return err
	}

	if strings.HasSuffix(origin, toType.Extension()) {
		return copyFile(origin, dst)
	}

	var cmd *exec.Cmd
	switch toType.Type() {
	case AudioM4A.Type():
		cmd = exec.Command("ffmpeg", "-i", origin, "-c:a", "aac", "-b:a", "192k", dst)
	case AudioWAV.Type():
		cmd = exec.Command("ffmpeg", "-i", origin, dst)
	default:
		return fmt.Errorf("unsupported audio type: %s", toType.Type())
	}

	return cmd.Run()
}

func copyFile(origin, dst string) error {
	sourceFile, err := os.Open(origin)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer func() {
		_ = sourceFile.Close()
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer func() {
		_ = dstFile.Close()
	}()

	_, err = io.Copy(dstFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	return nil
}

func checkDir(path string) error {
	base := filepath.Dir(path)
	if _, err := os.Stat(base); os.IsNotExist(err) {
		if err = os.MkdirAll(base, 0755); err != nil {
			return err
		}
	}

	return nil
}

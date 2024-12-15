package fsclients

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
)

func (c client) StoreAs(file *multipart.FileHeader, bucket, filename string, toType FileType) (string, error) {
	tempDst := c.makeTempPath(filepath.Ext(file.Filename)[1:])
	if err := saveFromFile(file, tempDst); err != nil {
		return "", err
	}

	defer func() {
		_ = os.Remove(tempDst)
	}()

	relativeFilePath := filepath.Join(bucket, fmt.Sprintf("%s.%s", filename, toType.Extension()))
	dst := filepath.Join(c.fsPath, relativeFilePath)

	// cutting corners here.
	// since it only accepts audioFiles there is no need to switch it.
	return relativeFilePath, convertAudioFile(tempDst, dst, toType)
}

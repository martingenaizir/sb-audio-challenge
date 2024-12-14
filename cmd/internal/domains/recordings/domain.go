package recordings

import (
	"context"
	"mime/multipart"
)

type Domain interface {
	RetrieveAs(ctx context.Context, filePath, format string) (string, error)
	StoreAs(file *multipart.FileHeader, dstFilename, dstExtension string) (string, error)
	RemoveFile(filePath string)
	ValidateFile(file *multipart.FileHeader) error
}

func Instance() Domain {
	// TODO
	return nil
}

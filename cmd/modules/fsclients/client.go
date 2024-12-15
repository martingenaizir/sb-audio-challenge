package fsclients

import (
	"context"
	"mime/multipart"
	"sync"
)

// of course, this should not be defined here.
const (
	uploadDir = "resources/uploads/"
	tempDir   = "resources/temp/"
)

type Client interface {
	StoreAs(file *multipart.FileHeader, bucket, filename string, toType FileType) (filePath string, err error)
	RetrieveAs(ctx context.Context, filePath string, asType FileType) (string, error)
	Remove(filePath string) error
}

type client struct {
	mu       *sync.Mutex
	fsPath   string
	tempPath string
}

// cutting corners here.
// this should be singleton and build by pool.

func Instance() Client {
	return &client{
		mu:       &sync.Mutex{},
		fsPath:   uploadDir,
		tempPath: tempDir,
	}
}

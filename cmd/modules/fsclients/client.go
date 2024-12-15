package fsclients

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"mime/multipart"
	"sync"
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

func Instance() Client {
	return &client{
		mu:       &sync.Mutex{},
		fsPath:   constants.UploadDir,
		tempPath: constants.TempDir,
	}
}

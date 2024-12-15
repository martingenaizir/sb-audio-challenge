package recordings

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/fsclients"
	"mime/multipart"
	"os"
)

type Domain interface {
	RetrieveAs(ctx context.Context, filePath, outFormat string) (string, string, error)
	StoreAs(file *multipart.FileHeader, dstFilename, dstExtension string) (StoredFile, error)
	RemoveFile(storedFile StoredFile)
	ValidateFile(file *multipart.FileHeader) error
}

type domain struct {
	withHistory bool
	fsClient    fsclients.Client
}

func Instance() Domain {
	return &domain{
		fsClient:    fsclients.Instance(),
		withHistory: os.Getenv(constants.WithRecordingsHistoryKey) == "true",
	}
}

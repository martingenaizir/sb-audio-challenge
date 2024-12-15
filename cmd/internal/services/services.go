package services

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/domains/phrases"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/domains/recordings"

	"mime/multipart"
)

type Services interface {
	GetUserRecordedPhrase(ctx context.Context, userID, phraseID int64, outputFormat string) (string, string, error)
	StoreUserRecordedPhrase(ctx context.Context, userID, phraseID int64, recordingFile *multipart.FileHeader) error
}

type services struct {
	phrases    phrases.Domain
	recordings recordings.Domain
}

func Instance() Services {
	return &services{
		phrases:    phrases.Instance(),
		recordings: recordings.Instance(),
	}
}

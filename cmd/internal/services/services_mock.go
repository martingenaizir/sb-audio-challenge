package services

import (
	"context"
	"mime/multipart"
)

type Mock struct {
	OnGetUserRecordedPhrase   func(ctx context.Context, userID, phraseID int64, outputFormat string) (string, error)
	OnStoreUserRecordedPhrase func(ctx context.Context, userID, phraseID int64, recordingFile *multipart.FileHeader) error
}

func (m Mock) GetUserRecordedPhrase(ctx context.Context, userID, phraseID int64, outputFormat string) (string, error) {
	if m.OnGetUserRecordedPhrase != nil {
		return m.OnGetUserRecordedPhrase(ctx, userID, phraseID, outputFormat)
	}

	return "", nil
}

func (m Mock) StoreUserRecordedPhrase(ctx context.Context, userID, phraseID int64, recordingFile *multipart.FileHeader) error {
	if m.OnStoreUserRecordedPhrase != nil {
		return m.OnStoreUserRecordedPhrase(ctx, userID, phraseID, recordingFile)
	}

	return nil
}

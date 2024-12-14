package services

import (
	"context"
	"fmt"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"mime/multipart"
)

// selected storing format.
const _targetFormat = "wav"

func (s services) StoreUserRecordedPhrase(ctx context.Context, userID, phraseID int64, recordingFile *multipart.FileHeader) error {
	if err := s.recordings.ValidateFile(recordingFile); err != nil {
		return err
	}

	practiceData, err := s.phrases.GetUserPracticeData(ctx, userID, phraseID)
	if err != nil {
		return err
	}

	if !practiceData.IsValidUser() || !practiceData.IsAccessiblePhrase() {
		return apierrors.NotFound("resource not found")
	}

	basename := fmt.Sprintf("practice_%d_%d", userID, phraseID)
	storedFile, sErr := s.recordings.StoreAs(recordingFile, basename, _targetFormat)
	if sErr != nil {
		return sErr
	}

	if err = s.phrases.StoreUserRecording(ctx, practiceData, storedFile); err != nil {
		go s.recordings.RemoveFile(storedFile)
		return err
	}

	return nil
}

package services

import (
	"context"
	"fmt"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"mime/multipart"
)

// selected storing format.
const (
	_targetFormat         = "wav"
	_practiceBaseFilename = "practice_%d_%d"
)

func (s services) StoreUserRecordedPhrase(ctx context.Context, userID, phraseID int64, file *multipart.FileHeader) error {
	if err := s.recordings.ValidateFile(file); err != nil {
		return err
	}

	practiceData, err := s.phrases.GetPracticeData(ctx, userID, phraseID)
	if err != nil {
		return err
	}

	if !practiceData.IsValidUser() || !practiceData.IsAccessiblePhrase() {
		return apierrors.NotFound("resource not found")
	}

	basename := fmt.Sprintf(_practiceBaseFilename, userID, phraseID)
	storedFile, sErr := s.recordings.StoreAs(file, basename, _targetFormat)
	if sErr != nil {
		return sErr
	}

	if err = s.phrases.StoreUserPractice(ctx, practiceData, storedFile.Path()); err != nil {
		go s.recordings.RemoveFile(storedFile)
		return err
	}

	return nil
}

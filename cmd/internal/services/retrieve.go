package services

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
)

func (s services) GetUserRecordedPhrase(ctx context.Context, userID, phraseID int64, outputFormat string) (string, error) {
	userPhraseData, err := s.phrases.GetUserPracticeData(ctx, userID, phraseID)
	if err != nil {
		return "", err
	}

	if !userPhraseData.IsValidUser() || !userPhraseData.IsAccessiblePhrase() || !userPhraseData.HasRecording() {
		return "", apierrors.NotFound("practice record not found")
	}

	return s.recordings.RetrieveAs(ctx, userPhraseData.RecordRelativePath(), outputFormat)
}

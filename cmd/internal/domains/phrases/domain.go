package phrases

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/dbclients"
)

type Domain interface {
	GetPracticeData(ctx context.Context, userID, phraseID int64) (UserPhraseData, error)
	StoreUserPractice(ctx context.Context, upd UserPhraseData, filePath string) error
}

type domain struct {
	dbRead      dbclients.Client
	dbWrite     dbclients.Client
	withHistory bool
}

func Instance() Domain {
	return &domain{
		dbRead:  dbclients.Get(constants.DBReadPool),
		dbWrite: dbclients.Get(constants.DBWritePool),
		// TODO from .env
		withHistory: constants.WithRecordingsHistory,
	}
}

package phrases

import (
	"context"
)

func (d *domain) GetPracticeData(ctx context.Context, userID, phraseID int64) (UserPhraseData, error) {
	data := UserPhraseData{}
	err := d.dbRead.Query(ctx, _getPracticeRecordData, func(f func(dest ...any) error) error {
		return f(&data.userID, &data.userLevel, &data.phraseID, &data.phraseLevel, &data.recordID, &data.recordRelativePath)
	}, phraseID, userID)

	return data, err
}

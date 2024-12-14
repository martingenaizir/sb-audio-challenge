package phrases

import (
	"context"
)

func (d *domain) StoreUserRecording(ctx context.Context, upd UserPhraseData, filePath string) error {
	if !d.withHistory && upd.HasRecording() {
		return d.updateRecording(ctx, upd.RecordID(), filePath)
	}

	return d.insertRecording(ctx, upd.userID.Int64, upd.phraseID.Int64, filePath)
}

func (d *domain) updateRecording(ctx context.Context, recordID int64, path string) error {
	tx, err := d.dbWrite.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	if _, err = tx.Exec(ctx, _updatePracticeRecordVersion, path, recordID); err != nil {
		return err
	}

	return tx.Commit()
}

func (d *domain) insertRecording(ctx context.Context, userID, phraseID int64, path string) error {
	tx, err := d.dbWrite.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	if _, err = tx.Exec(ctx, _insertPracticeRecordVersion, userID, phraseID, path); err != nil {
		return err
	}

	return tx.Commit()
}

package phrases

import "database/sql"

// UserPhraseData from DB.
type UserPhraseData struct {
	userID             sql.NullInt64
	userLevel          sql.NullInt64
	phraseID           sql.NullInt64
	phraseLevel        sql.NullInt64
	recordID           sql.NullInt64
	recordFileName     sql.NullString
	recordRelativePath sql.NullString
}

func (d UserPhraseData) IsValidUser() bool {
	return d.userID.Valid && d.userID.Int64 > 0
}

func (d UserPhraseData) IsAccessiblePhrase() bool {
	return d.phraseID.Valid && d.phraseID.Int64 > 0 &&
		d.phraseLevel.Valid &&
		d.userLevel.Valid &&
		d.phraseLevel.Int64 <= d.userLevel.Int64
}

func (d UserPhraseData) HasRecording() bool {
	return d.recordID.Valid && d.recordID.Int64 > 0
}

func (d UserPhraseData) RecordPath() string {
	return d.recordRelativePath.String
}

func (d UserPhraseData) RecordID() int64 {
	return d.recordID.Int64
}

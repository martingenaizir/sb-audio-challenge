package phrases

// If user exists, this query can tell:
// if the phrase exist: phrases.phrase_id > 0.
// if user has access to the phrase: phrases.level <= user_level
// if the user has already practiced it: user_phrases.id > 0.
//
// args: phrase_id, user_id
const _getPracticeRecordData = `SELECT 
	users.user_id
	, users.user_level 
	, phrases.phrase_id
    , phrases.level
	, user_phrases.id
	, user_phrases.filepath
FROM users  
LEFT JOIN phrases ON 
	phrases.phrase_id = ? 
	AND phrases.level < user_level 
	AND phrases.active = 1
left join user_phrases ON 
	user_phrases.user_id = users.user_id 
	AND user_phrases.phrase_id = phrases.phrase_id 
WHERE 
	users.user_id = ? 
	AND users.active
ORDER BY user_phrases.id DESC
LIMIT 1`

// args: filepath, id.
const _updatePracticeRecordVersion = `UPDATE user_phrases SET filepath = ? WHERE id = ?`

// args: user_id, phrase_id, filepath.
const _insertPracticeRecordVersion = `INSERT INTO user_phrases(user_id, phrase_id, filepath) VALUES (?,?,?)`

package constants

const GinModeKey = "GIN_MODE"

// app.
const (
	AppPortKey                     = "APP_PORT"
	LogLevelKey                    = "APP_LOG_LEVEL"
	LogLevelDef                    = "info"
	WithRecordingsHistoryKey       = "APP_WITH_RECORDING_HISTORY"
	MaxAudioFileSizeBytes    int64 = 5000000
)

// DB.
const (
	DBReadPool  = "read"
	DBWritePool = "write"
	DBHostKey   = "MYSQL_HOST"
	DBPortKey   = "MYSQL_PORT"
	DBNameKey   = "MYSQL_DATABASE"
	DBUserKey   = "MYSQL_USER"
	DBPassKey   = "MYSQL_PASSWORD"
	DBQueryKey  = "DB_CON_QUERY"
)

var DBPools = []string{DBReadPool, DBWritePool}

// params.
const (
	UserIDParamKey      = "user_id"
	PhraseIDParamKey    = "phrase_id"
	AudioFileFormKey    = "audio_file"
	AudioFormatParamKey = "audio_format"
)

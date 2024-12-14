package constants

// Files.
var (
	UploadDir = "resources/uploads/"
	TempDir   = "resources/temp/"
)

const GinModeKey = "GIN_MODE"

// App.
const (
	AppPortKey            = "APP_PORT"
	LogLevelKey           = "APP_LOG_LEVEL"
	LogLevelDef           = "info"
	WithRecordingsHistory = false
)

// DB.
const (
	DBReadPool  = "read"
	DBWritePool = "write"
	DBHostKey   = "DB_HOST"
	DBPortKey   = "DB_PORT"
	DBNameKey   = "DB_NAME"
	DBUserKey   = "DB_USER"
	DBPassKey   = "DB_PASS"
)

var DBPools = []string{DBReadPool, DBWritePool}

// Params.
const (
	UserIDParamKey      = "user_id"
	PhraseIDParamKey    = "phrase_id"
	AudioFileFormKey    = "audio_file"
	AudioFormatParamKey = "audio_format"
)

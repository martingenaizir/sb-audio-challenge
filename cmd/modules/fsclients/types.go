package fsclients

import "strings"

type FileType interface {
	Type() string
	Extension() string
}

type AudioTypes struct {
	mimeType  string
	extension string
}

func (at AudioTypes) Type() string {
	return at.mimeType
}

func (at AudioTypes) Extension() string {
	return at.extension
}

var (
	WAV = AudioTypes{
		mimeType:  "audio/wav",
		extension: "wav",
	}

	M4A = AudioTypes{
		mimeType:  "audio/mp4",
		extension: "m4a",
	}

	// TODO more types?
)

func IsSameType(originPath string, fileType FileType) bool {
	return strings.HasSuffix(originPath, fileType.Extension())
}

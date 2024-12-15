package fsclients

import (
	"slices"
	"strings"
)

type FileType interface {
	Type() string
	Extension() string
	Aliases() []string
}

type AudioType struct {
	mimeType  string
	extension string
	aliases   []string
}

func (at AudioType) Type() string {
	return at.mimeType
}

func (at AudioType) Extension() string {
	return at.extension
}

func (at AudioType) Aliases() []string {
	return at.aliases
}

var (
	AudioWAV = AudioType{
		mimeType:  "audio/wav",
		extension: "wav",
		aliases:   []string{},
	}

	AudioM4A = AudioType{
		mimeType:  "audio/mp4",
		extension: "m4a",
		aliases:   []string{"mp4"},
	}
)

var _supportedTypes = []FileType{AudioM4A, AudioWAV}

func IsSameType(originPath string, fileType FileType) bool {
	return strings.HasSuffix(originPath, fileType.Extension())
}

func CastType(t, filename string) (FileType, bool) {
	for _, st := range _supportedTypes {
		if st.Extension() == t || st.Type() == t || slices.Contains(st.Aliases(), t) {
			return st, true
		}

		// cutting corners here
		// this should an actual validation
		if t == "application/octet-stream" && strings.HasSuffix(filename, st.Extension()) {
			return st, true
		}
	}

	return AudioWAV, false
}

package recordings

import (
	"fmt"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/fsclients"

	"mime/multipart"
)

const (
	_contentTypeHeaderKey = "Content-Type"
	_toMB                 = 1024 * 1024
)

func (d *domain) ValidateFile(file *multipart.FileHeader) error {
	if file == nil {
		return apierrors.BadRequestError("missing or invalid file")
	}

	if file.Size > constants.MaxAudioFileSizeBytes {
		return apierrors.ContentTooLargeError(fmt.Sprintf("the file cannot exceed %.2f MB",
			float64(constants.MaxAudioFileSizeBytes)/_toMB))
	}

	if _, ok := fsclients.CastType(file.Header.Get(_contentTypeHeaderKey), file.Filename); !ok {
		return apierrors.BadRequestError("unsupported file type")
	}

	return nil
}

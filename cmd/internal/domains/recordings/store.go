package recordings

import (
	"fmt"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/fsclients"

	"mime/multipart"
	"time"
)

const _bucket = "recordings"

func (d *domain) StoreAs(file *multipart.FileHeader, basename, extension string) (StoredFile, error) {
	storeWithType, ok := fsclients.CastType(extension, basename)
	if !ok {
		return StoredFile{}, apierrors.BadRequestError("invalid target format")
	}

	filename := basename
	if d.withHistory {
		filename = fmt.Sprintf("%d-%s", time.Now().UnixMilli(), basename)
	}

	path, err := d.storage.StoreAs(file, _bucket, filename, storeWithType)
	if err != nil {
		return StoredFile{}, err
	}

	return StoredFile{
		name: filename,
		path: path,
	}, nil
}

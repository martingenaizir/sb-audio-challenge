package recordings

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/fsclients"
)

func (d *domain) RetrieveAs(ctx context.Context, filePath, outFormat string) (string, error) {
	wantedFormat, ok := fsclients.CastType(outFormat, "")
	if !ok {
		return "", apierrors.BadRequestError("invalid format")
	}

	return d.storage.RetrieveAs(ctx, filePath, wantedFormat)
}

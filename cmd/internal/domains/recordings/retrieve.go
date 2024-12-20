package recordings

import (
	"context"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/fsclients"
)

func (d *domain) RetrieveAs(ctx context.Context, filePath, outFormat string) (string, string, error) {
	audioType, ok := fsclients.NewAudioType(outFormat, "")
	if !ok {
		return "", "", apierrors.BadRequestError("invalid format")
	}

	outPath, err := d.fsClient.RetrieveAs(ctx, filePath, audioType)
	return outPath, audioType.Type(), err
}

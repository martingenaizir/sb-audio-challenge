package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"mime/multipart"
	"net/http"
)

func (c Controller) StoreUserPracticePhrase(ctx *gin.Context) {
	file, fileErr := formFile(ctx, constants.AudioFileFormKey)
	if fileErr != nil {
		_ = ctx.Error(apierrors.BadRequestError("missing or invalid audio file"))
	}

	if err := c.services.StoreUserRecordedPhrase(
		ctx.Request.Context(),
		ctx.GetInt64(constants.UserIDParamKey),
		ctx.GetInt64(constants.PhraseIDParamKey),
		file,
	); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

var formFile = func(ctx *gin.Context, name string) (*multipart.FileHeader, error) {
	return ctx.FormFile(name)
}

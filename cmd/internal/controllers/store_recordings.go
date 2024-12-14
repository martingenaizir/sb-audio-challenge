package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"net/http"
)

func (c Controller) StoreUserPracticePhrase(ctx *gin.Context) {
	_, fileErr := ctx.FormFile(constants.AudioFileFormKey)
	if fileErr != nil {
		_ = ctx.Error(apierrors.BadRequestError("missing or invalid audio file"))
		return
	}

	// TODO

	ctx.Status(http.StatusServiceUnavailable)
}

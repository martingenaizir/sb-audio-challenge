package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"net/http"
	"strings"
)

const _supportedFormat = "m4a"

func (c Controller) GetRecordedPracticePhrase(ctx *gin.Context) {
	toFormat := ctx.Param(constants.AudioFormatParamKey)

	if !strings.EqualFold(toFormat, _supportedFormat) {
		_ = ctx.Error(apierrors.BadRequestError("unsupported audio format"))
		return
	}

	// TODO

	ctx.Status(http.StatusServiceUnavailable)
}

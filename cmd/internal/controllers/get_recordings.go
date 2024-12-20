package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"path/filepath"
	"slices"
)

var _supportedFormats = []string{"wav", "mp3", "mp4", "m4a"}

func (c Controller) GetRecordedPracticePhrase(ctx *gin.Context) {
	toFormat := ctx.Param(constants.AudioFormatParamKey)
	if !slices.Contains(_supportedFormats, toFormat) {
		_ = ctx.Error(apierrors.BadRequestError("unsupported audio format"))
		return
	}

	userID := ctx.GetInt64(constants.UserIDParamKey)
	phraseID := ctx.GetInt64(constants.PhraseIDParamKey)
	filePath, mimeType, err := c.services.GetUserRecordedPhrase(ctx.Request.Context(), userID, phraseID, toFormat)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	filename := fmt.Sprintf("audio_%d_%d%s", userID, phraseID, filepath.Ext(filePath))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", mimeType)
	ctx.Header("Content-Length", "0")
	serveFile(ctx, filePath)
}

var serveFile = func(ctx *gin.Context, filePath string) {
	ctx.File(filePath)
}

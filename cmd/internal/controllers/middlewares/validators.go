package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"regexp"
	"strconv"
)

var _positiveIntRegex = regexp.MustCompile(`^[1-9][0-9]*$`)

func PhrasesPathValidation(ctx *gin.Context) {
	defer func() {
		if len(ctx.Errors) > 0 {
			ctx.Abort()
		}
	}()

	userStrID := ctx.Param(constants.UserIDParamKey)
	if !_positiveIntRegex.MatchString(userStrID) {
		_ = ctx.Error(apierrors.BadRequestError("invalid user id"))
		return
	}

	phraseStrID := ctx.Param(constants.PhraseIDParamKey)
	if !_positiveIntRegex.MatchString(phraseStrID) {
		_ = ctx.Error(apierrors.BadRequestError("invalid phrase id"))
		return
	}

	ctx.Set(constants.UserIDParamKey, asInt64(userStrID))
	ctx.Set(constants.PhraseIDParamKey, asInt64(phraseStrID))
}

func asInt64(value string) int64 {
	intValue, _ := strconv.ParseInt(value, 10, 64)
	return intValue
}

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"strconv"
)

type pathIDs struct {
	UserID, PhraseID int64 `validate:"gt=0"`
}

var val = validator.New()

func PhrasesPathValidation(ctx *gin.Context) {
	defer func() {
		if len(ctx.Errors) > 0 {
			ctx.Abort()
		}
	}()
	
	data := pathIDs{
		UserID:   asInt64(ctx.Param(constants.UserIDParamKey)),
		PhraseID: asInt64(ctx.Param(constants.PhraseIDParamKey)),
	}

	if err := val.Struct(data); err != nil {
		_ = ctx.Error(apierrors.BadRequestError(err.Error()))
		return
	}

	ctx.Set(constants.UserIDParamKey, data.UserID)
	ctx.Set(constants.PhraseIDParamKey, data.PhraseID)
}

func asInt64(value string) int64 {
	intValue, _ := strconv.ParseInt(value, 10, 64)
	return intValue
}

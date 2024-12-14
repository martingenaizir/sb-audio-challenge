package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/logger"
	"net/http"
)

const _errorTemplate = "%s %s"

func ApiErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors.Last()

		logger.Error(err, _errorTemplate, ctx.Request.Method, ctx.Request.URL.Redacted())

		var apiErr apierrors.ApiError
		if errors.As(err, &apiErr) {
			ctx.JSON(apiErr.Code(), apiErr)
			return
		}

		// the err.Error() is just for dev / test.
		// prod scopes shouldn't expose unchecked info.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
	}
}

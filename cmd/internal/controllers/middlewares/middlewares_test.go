package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/apierrors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApiErrorHandler(t *testing.T) {
	type args struct {
		prepare func() (*gin.Context, *httptest.ResponseRecorder)
	}
	tests := []struct {
		name string
		args args
		code int
	}{
		{
			name: "known error",
			args: args{
				prepare: func() (*gin.Context, *httptest.ResponseRecorder) {
					r, _ := http.NewRequest(http.MethodPost, "/", nil)
					ctx, rr := _mockCtx(r)
					_ = ctx.Error(apierrors.BadRequestError("bad request"))
					return ctx, rr
				},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "unknown error",
			args: args{
				prepare: func() (*gin.Context, *httptest.ResponseRecorder) {
					r, _ := http.NewRequest(http.MethodPost, "/", nil)
					ctx, rr := _mockCtx(r)
					_ = ctx.Error(errors.New("mocked error"))
					return ctx, rr
				},
			},
			code: http.StatusInternalServerError,
		},
		{
			name: "no errors",
			args: args{
				prepare: func() (*gin.Context, *httptest.ResponseRecorder) {
					r, _ := http.NewRequest(http.MethodPost, "/", nil)
					return _mockCtx(r)
				},
			},
			code: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, rr := tt.args.prepare()
			ApiErrorHandler(ctx)

			if rr.Code != tt.code {
				t.Errorf("ApiErrorHandler() = %v, want %v", rr.Code, tt.code)
			}
		})
	}
}

func TestPhrasesPathValidation(t *testing.T) {
	goodID := "1"
	badID := "-1"

	type args struct {
		prepare func() *gin.Context
	}
	tests := []struct {
		name      string
		args      args
		withError bool
		errPart   string
	}{
		{
			name: "with invalid data",
			args: args{
				prepare: func() *gin.Context {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					ctx, _ := _mockCtx(r)
					ctx.AddParam(constants.UserIDParamKey, "lorem")
					ctx.AddParam(constants.PhraseIDParamKey, goodID)
					return ctx
				},
			},
			withError: true,
			errPart:   "invalid user id",
		},
		{
			name: "invalid user_id",
			args: args{
				prepare: func() *gin.Context {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					ctx, _ := _mockCtx(r)
					ctx.AddParam(constants.UserIDParamKey, badID)
					ctx.AddParam(constants.PhraseIDParamKey, goodID)
					return ctx
				},
			},
			withError: true,
			errPart:   "invalid user id",
		},
		{
			name: "invalid phrase_id",
			args: args{
				prepare: func() *gin.Context {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					ctx, _ := _mockCtx(r)
					ctx.AddParam(constants.UserIDParamKey, goodID)
					ctx.AddParam(constants.PhraseIDParamKey, badID)
					return ctx
				},
			},
			withError: true,
			errPart:   "invalid phrase id",
		},
		{
			name: "no errors",
			args: args{
				prepare: func() *gin.Context {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					ctx, _ := _mockCtx(r)
					ctx.AddParam(constants.UserIDParamKey, goodID)
					ctx.AddParam(constants.PhraseIDParamKey, goodID)
					return ctx
				},
			},
			withError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.args.prepare()
			PhrasesPathValidation(ctx)
			if tt.withError && !strings.Contains(ctx.Errors.String(), tt.errPart) {
				t.Errorf("PhrasesPathValidation() = %v, want %v", ctx.Errors.String(), tt.errPart)
				return
			}

			if !tt.withError && len(ctx.Errors) > 0 {
				t.Errorf("PhrasesPathValidation() = %v, want nil", ctx.Errors.String())
			}
		})
	}
}

func _mockCtx(r *http.Request) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Request = r
	return ctx, rr
}

package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/services"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_GetRecordedPracticePhrase(t *testing.T) {
	someID := int64(1)

	type fields struct {
		services services.Services
	}
	type args struct {
		prepare func() *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		withErr bool
	}{
		{
			name: "unsupported format",
			fields: fields{
				services: services.Mock{},
			},
			args: args{
				prepare: _mockCtx,
			},
			withErr: true,
		},
		{
			name: "get recording error",
			fields: fields{
				services: services.Mock{OnGetUserRecordedPhrase: func(_ context.Context, _, _ int64, _ string) (string, error) {
					return "", errors.New("error")
				}},
			},
			args: args{
				prepare: func() *gin.Context {
					ctx := _mockCtx()
					ctx.Set(constants.UserIDParamKey, someID)
					ctx.Set(constants.PhraseIDParamKey, someID)
					ctx.AddParam(constants.AudioFormatParamKey, "m4a")
					return ctx
				},
			},
			withErr: true,
		},
		{
			name: "serve file success",
			fields: fields{
				services: services.Mock{},
			},
			args: args{
				prepare: func() *gin.Context {
					serveFile = func(ctx *gin.Context, filePath string) {

					}

					ctx := _mockCtx()
					ctx.Set(constants.UserIDParamKey, someID)
					ctx.Set(constants.PhraseIDParamKey, someID)
					ctx.AddParam(constants.AudioFormatParamKey, "m4a")
					return ctx
				},
			},
			withErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				services: tt.fields.services,
			}

			ctx := tt.args.prepare()
			c.GetRecordedPracticePhrase(ctx)

			if tt.withErr && len(ctx.Errors) == 0 {
				t.Error("Controller.GetRecordedPracticePhrase() should have returned an error")
			}
		})
	}
}

func TestController_StoreUserPracticePhrase(t *testing.T) {
	type fields struct {
		services services.Services
	}
	type args struct {
		prepare func() *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		withErr bool
	}{
		{
			name: "file error",
			fields: fields{
				services: services.Mock{},
			},
			args: args{
				prepare: func() *gin.Context {
					formFile = func(_ *gin.Context, _ string) (*multipart.FileHeader, error) {
						return nil, errors.New("file error")
					}
					return _mockCtx()
				},
			},
			withErr: true,
		},
		{
			name: "store error",
			fields: fields{
				services: services.Mock{OnStoreUserRecordedPhrase: func(_ context.Context, _, _ int64, _ *multipart.FileHeader) error {
					return errors.New("store error")
				}},
			},
			args: args{
				prepare: func() *gin.Context {
					formFile = func(_ *gin.Context, _ string) (*multipart.FileHeader, error) {
						return &multipart.FileHeader{}, nil
					}
					return _mockCtx()
				},
			},
			withErr: true,
		},
		{
			name: "store success",
			fields: fields{
				services: services.Mock{},
			},
			args: args{
				prepare: func() *gin.Context {
					formFile = func(_ *gin.Context, _ string) (*multipart.FileHeader, error) {
						return &multipart.FileHeader{}, nil
					}
					return _mockCtx()
				},
			},
			withErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				services: tt.fields.services,
			}

			ctx := tt.args.prepare()
			c.StoreUserPracticePhrase(ctx)

			if tt.withErr && len(ctx.Errors) == 0 {
				t.Error("Controller.GetRecordedPracticePhrase() should have returned an error")
			}
		})
	}
}

func _mockCtx() *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	return ctx
}

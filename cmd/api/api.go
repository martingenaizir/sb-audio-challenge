package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/internal/controllers"
	mw "github.com/martingenaizir/sb-audio-challenge/cmd/internal/controllers/middlewares"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules"
	"net/http"
	"os"
)

func newApplication() error {
	if err := modules.Configure(); err != nil {
		return err
	}

	gin.SetMode(gin.DebugMode)
	g := gin.Default()

	if err := g.SetTrustedProxies(nil); err != nil {
		return err
	}

	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	c := controllers.Instance()

	g.Use(mw.ApiErrorHandler, mw.PhrasesPathValidation)
	up := g.Group("/audio/user/:user_id/phrase/:phrase_id")
	{
		up.POST("", c.StoreUserPracticePhrase)
		up.GET("/:audio_format", c.GetRecordedPracticePhrase)
	}

	return g.Run(fmt.Sprintf(":%s", os.Getenv(constants.AppPortKey)))
}

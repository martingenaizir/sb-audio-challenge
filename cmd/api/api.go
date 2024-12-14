package main

import (
	"github.com/gin-gonic/gin"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules"
	"net/http"
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

	return g.Run(":8080")
}

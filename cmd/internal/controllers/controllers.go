package controllers

import "github.com/martingenaizir/sb-audio-challenge/cmd/internal/services"

type Controller struct {
	services services.Services
}

func Instance() *Controller {
	return &Controller{
		services: services.Instance(),
	}
}

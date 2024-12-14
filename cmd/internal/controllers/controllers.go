package controllers

type Controller struct {
}

func Instance() *Controller {
	return &Controller{}
}

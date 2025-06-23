package routers

import (
	"go_project_structure_be/handler"
	middlewares "go_project_structure_be/middleware"

	"github.com/gofiber/fiber/v2"
)

type ApiRegistry interface {
	// Authorization()
	UsersAPi()
}

type apiRegistry struct {
	hand *handler.Handler
	app  *fiber.App
	m    middlewares.MiddleWare
}

func NewapiRegistry(hand *handler.Handler, app *fiber.App, m middlewares.MiddleWare) ApiRegistry {
	return &apiRegistry{hand: hand, app: app, m: m}
}

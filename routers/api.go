package routers

import (
	"go_ci/handler"
	middlewares "go_ci/middleware"

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

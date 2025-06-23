package middlewares

import (
	"go_project_structure_be/configurations"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/utils"
)

type MiddleWare interface {
	Static() fiber.Router
	Csrf() fiber.Router
	Logger() fiber.Router
	Cors() fiber.Router
	AuthRequired() fiber.Handler
}

type middleWare struct {
	app   *fiber.App
	confs *configurations.Configs
}

func NewMiddleWare(app *fiber.App, conf configurations.Configs) MiddleWare {
	return &middleWare{app: app, confs: &conf}
}

func (m *middleWare) Static() fiber.Router {
	configStatic := fiber.Static{
		Index:         "index.html",
		CacheDuration: 5 * time.Second,
	}
	return m.app.Static("/", "./static/public", configStatic)

}

func (m *middleWare) Csrf() fiber.Router {

	configCSRF := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "_csrf_",
	}
	return m.app.Use(csrf.New(configCSRF))
}

func (m *middleWare) Logger() fiber.Router {
	return m.app.Use(logger.New())
}

func (m *middleWare) Cors() fiber.Router {
	configCores := cors.Config{
		AllowOrigins:     m.confs.Conf.Cors.AllowOrigins,
		AllowCredentials: m.confs.Conf.Cors.AllowCredentials,
		AllowHeaders:     m.confs.Conf.Cors.AllowHeaders,
		AllowMethods:     m.confs.Conf.Cors.AllowMethods,
	}
	return m.app.Use(cors.New(configCores))
}

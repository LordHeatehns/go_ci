package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func (m *middleWare) AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(m.confs.ConfEnv.Secret_token),
			JWTAlg: jwtware.HS256,
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": fiber.ErrUnauthorized.Message,
				"code":    fiber.StatusUnauthorized,
				"success": false,
			})
		},
	})
}

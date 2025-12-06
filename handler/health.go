package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"code":    fiber.StatusOK,
	})
}

func (h *Handler) Ready(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"code":    fiber.StatusOK,
	})
}

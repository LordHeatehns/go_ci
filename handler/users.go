package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {

	datas, err := h.sv.GetUsers()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    datas,
		"code":    fiber.StatusOK,
	})
}

func (h *Handler) GetUsersTest(c *fiber.Ctx) error {

	datas := []interface{}{
		map[string]interface{}{"id": 1, "name": "John"},
		map[string]interface{}{"id": 2, "name": "Jane"},
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    datas,
		"code":    fiber.StatusOK,
	})
}

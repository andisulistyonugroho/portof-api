package controller

import (
	"github.com/gofiber/fiber/v3"
)

func (cdh *DuckHandler) ListCategory(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "ok",
	})
}

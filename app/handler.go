package app

import (
	"net/http"
	"photoApi/services"

	"github.com/gofiber/fiber/v2"
)

type PhotoHandler struct {
	Service services.Service
}

func (h PhotoHandler) GetAllPhoto(c *fiber.Ctx) error {
	result, err := h.Service.PhotoGetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(result)
}

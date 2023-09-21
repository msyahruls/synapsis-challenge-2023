package controller

import (
	"synapsis-challange/model/web"

	"github.com/gofiber/fiber/v2"
)

func UnauthorizeReturn(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
		Code:    fiber.StatusUnauthorized,
		Status:  false,
		Message: message,
	})
}

package exception

import (
	"synapsis-challange/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "bad request",
			Data:    err.Error(),
		})
	}
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(web.WebResponse{
			Code:    e.Code,
			Status:  false,
			Message: e.Message,
			Data:    err.Error(),
		})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
		Code:    fiber.StatusInternalServerError,
		Status:  false,
		Message: "internal server error",
		Data:    err.Error(),
	})
}

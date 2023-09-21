package controller

import (
	"synapsis-challange/middleware"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type LogControllerImpl struct {
	LogService service.LogService
}

type LogController interface {
	NewLogRouter(app *fiber.App)
}

func NewLogController(logService service.LogService) LogController {
	return &LogControllerImpl{
		LogService: logService,
	}
}

func (controller *LogControllerImpl) NewLogRouter(app *fiber.App) {
	log := app.Group("/logs")
	log.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	log.Get("/", controller.FindAll)
}

func (controller *LogControllerImpl) FindAll(ctx *fiber.Ctx) error {
	if isadmin := middleware.IsAdmin(ctx); !isadmin {
		return UnauthorizeReturn(ctx, "unathorize")
	}

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	log := controller.LogService.FindAll(page, limit)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    log,
	})
}

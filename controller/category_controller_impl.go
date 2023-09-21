package controller

import (
	"fmt"

	"synapsis-challange/helper"
	middlewares "synapsis-challange/middleware"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
	LogService      service.LogService
}

type CategoryController interface {
	NewCategoryRouter(app *fiber.App)
}

func NewCategoryController(categoryService service.CategoryService, logService service.LogService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
		LogService:      logService,
	}
}

func (controller *CategoryControllerImpl) NewCategoryRouter(app *fiber.App) {
	category := app.Group("/categories")
	category.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	// //Restricted (JUST FOR ADMIN)
	category.Use(middlewares.IsAuthenticated)
	category.Post("/", controller.Create)
	category.Get("/:categoryId", controller.FindById)
	category.Get("/", controller.FindAll)
	category.Put("/:categoryId", controller.Update)
	category.Delete("/:categoryId", controller.Delete)
}

func (controller *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can create category")
	// }

	var request web.CategoryCreateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(request)

	action := fmt.Sprintf("create category %s", categoryResponse.Name)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindById(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	categoryResponse := controller.CategoryService.FindById(categoryId)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	categoryResponse := controller.CategoryService.FindAll(name)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can update category")
	// }

	var request web.CategoryUpdateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	categoryId := ctx.Params("categoryId")

	categoryResponse := controller.CategoryService.Update(categoryId, request)

	action := fmt.Sprintf("update category %s", categoryResponse.Name)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can delete category")
	// }

	categoryId := ctx.Params("categoryId")

	controller.CategoryService.Delete(categoryId)

	action := fmt.Sprintf("delete category %s", categoryId)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    nil,
	})
}

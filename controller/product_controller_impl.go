package controller

import (
	"fmt"

	"synapsis-challange/helper"
	middlewares "synapsis-challange/middleware"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
	LogService     service.LogService
}

type ProductController interface {
	NewProductRouter(app *fiber.App)
}

func NewProductController(productService service.ProductService, logService service.LogService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
		LogService:     logService,
	}
}

func (controller *ProductControllerImpl) NewProductRouter(app *fiber.App) {
	product := app.Group("/products")
	product.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	// //Restricted (JUST FOR ADMIN)
	product.Use(middlewares.IsAuthenticated)
	product.Post("/", controller.Create)
	product.Get("/:productId", controller.FindById)
	product.Get("/", controller.FindAll)
	product.Put("/:productId", controller.Update)
	product.Delete("/:productId", controller.Delete)
}

func (controller *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	var request web.ProductCreateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.Create(request)

	action := fmt.Sprintf("create product %s", productResponse.Name)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindById(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	productResponse := controller.ProductService.FindById(productId)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) FindAll(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	category := ctx.Query("category")
	productResponse := controller.ProductService.FindAll(name, category)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	var request web.ProductUpdateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	productId := ctx.Params("productId")

	productResponse := controller.ProductService.Update(productId, request)

	action := fmt.Sprintf("update product %s", productResponse.Name)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    productResponse,
	})
}

func (controller *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	controller.ProductService.Delete(productId)

	action := fmt.Sprintf("delete product %s", productId)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    nil,
	})
}

package controller

import (
	"fmt"

	"synapsis-challange/helper"
	middlewares "synapsis-challange/middleware"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type CartControllerImpl struct {
	CartService service.CartService
	LogService  service.LogService
}

type CartController interface {
	NewCartRouter(app *fiber.App)
}

func NewCartController(cartService service.CartService, logService service.LogService) CartController {
	return &CartControllerImpl{
		CartService: cartService,
		LogService:  logService,
	}
}

func (controller *CartControllerImpl) NewCartRouter(app *fiber.App) {
	cart := app.Group("/cart")
	cart.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	// //Restricted (JUST FOR ADMIN)
	cart.Use(middlewares.IsAuthenticated)
	cart.Post("/", controller.Create)
	cart.Get("/:cartId", controller.FindById)
	cart.Get("/", controller.FindAll)
	cart.Put("/:cartId", controller.Update)
	cart.Delete("/:cartId", controller.Delete)
}

func (controller *CartControllerImpl) Create(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can create cart")
	// }

	var request web.CartCreateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)
	// fmt.Println(request)

	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))

	request.UserID = actor

	fmt.Println(request)

	cartResponse := controller.CartService.Create(request, actor)

	action := fmt.Sprintf("add %s to cart", cartResponse.ProductID)
	// actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    cartResponse,
	})
}

func (controller *CartControllerImpl) FindById(ctx *fiber.Ctx) error {
	cartId := ctx.Params("cartId")

	cartResponse := controller.CartService.FindById(cartId)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    cartResponse,
	})
}

func (controller *CartControllerImpl) FindAll(ctx *fiber.Ctx) error {
	// name := ctx.Query("name")
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	cartResponse := controller.CartService.FindAll(actor, "")
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    cartResponse,
	})
}

func (controller *CartControllerImpl) Update(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can update cart")
	// }

	var request web.CartUpdateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))

	request.UserID = actor

	cartId := ctx.Params("cartId")

	cartResponse := controller.CartService.Update(cartId, request)

	action := fmt.Sprintf("update %s from cart", cartResponse.ProductID)
	// actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    cartResponse,
	})
}

func (controller *CartControllerImpl) Delete(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can delete cart")
	// }

	cartId := ctx.Params("cartId")

	controller.CartService.Delete(cartId)

	action := fmt.Sprintf("delete list cart %s", cartId)
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    nil,
	})
}

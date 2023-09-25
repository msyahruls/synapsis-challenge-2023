package controller

import (
	"fmt"
	"synapsis-challange/helper"
	middlewares "synapsis-challange/middleware"
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
	CartService        service.CartService
	LogService         service.LogService
}

type TransactionController interface {
	NewTransactionRouter(app *fiber.App)
}

func NewTransactionController(transactionService service.TransactionService, cartService service.CartService, logService service.LogService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: transactionService,
		CartService:        cartService,
		LogService:         logService,
	}
}

func (controller *TransactionControllerImpl) NewTransactionRouter(app *fiber.App) {
	transaction := app.Group("/transactions")
	transaction.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	// //Restricted (JUST FOR ADMIN)
	transaction.Use(middlewares.IsAuthenticated)
	transaction.Post("/checkout", controller.Create)
	transaction.Get("/:transactionId", controller.FindById)
	transaction.Get("/", controller.FindAll)
	// transaction.Put("/:transactionId", controller.Update)
	// transaction.Delete("/:transactionId", controller.Delete)
}

func (controller *TransactionControllerImpl) Create(ctx *fiber.Ctx) error {
	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
	// 	return UnauthorizeReturn(ctx, "only admin can create transaction")
	// }

	var request web.TransactionCreateRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	request.UserID = actor

	cartResponse := controller.CartService.FindAll(actor, "")

	sumTotal := 0
	var targetArray []domain.TransactionDetails
	for _, cart := range cartResponse {
		target := domain.TransactionDetails{
			ProductID: cart.ProductID,
			Qty:       cart.Qty,
			Total:     cart.Total,
		}
		targetArray = append(targetArray, target)
		sumTotal += target.Total
	}
	request.Transactions = targetArray
	request.SumTotal = sumTotal
	request.Payback = request.Payment - sumTotal

	fmt.Println(request)

	transactionResponse := controller.TransactionService.Create(request, actor)

	action := fmt.Sprintf("add %s to transaction", transactionResponse.UserID)
	// actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) FindById(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")

	transactionResponse := controller.TransactionService.FindById(transactionId)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    transactionResponse,
	})
}

func (controller *TransactionControllerImpl) FindAll(ctx *fiber.Ctx) error {
	// name := ctx.Query("name")
	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
	transactionResponse := controller.TransactionService.FindAll(actor)
	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    transactionResponse,
	})
}

// func (controller *TransactionControllerImpl) Update(ctx *fiber.Ctx) error {
// 	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
// 	// 	return UnauthorizeReturn(ctx, "only admin can update transaction")
// 	// }

// 	var request web.TransactionUpdateRequest
// 	err := ctx.BodyParser(&request)
// 	helper.PanicIfError(err)
// 	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))

// 	request.UserID = actor

// 	transactionId := ctx.Params("transactionId")

// 	transactionResponse := controller.TransactionService.Update(transactionId, request)

// 	action := fmt.Sprintf("update %s from transaction", transactionResponse.ProductID)
// 	// actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
// 	controller.LogService.Create(actor, action)

// 	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
// 		Code:    fiber.StatusOK,
// 		Status:  true,
// 		Message: "success",
// 		Data:    transactionResponse,
// 	})
// }

// func (controller *TransactionControllerImpl) Delete(ctx *fiber.Ctx) error {
// 	// if isAdmin := middlewares.IsAdmin(ctx); !isAdmin {
// 	// 	return UnauthorizeReturn(ctx, "only admin can delete transaction")
// 	// }

// 	transactionId := ctx.Params("transactionId")

// 	controller.TransactionService.Delete(transactionId)

// 	action := fmt.Sprintf("delete list transaction %s", transactionId)
// 	actor, _, _ := helper.ParseJwt(ctx.Cookies("token"))
// 	controller.LogService.Create(actor, action)

// 	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
// 		Code:    fiber.StatusOK,
// 		Status:  true,
// 		Message: "success",
// 		Data:    nil,
// 	})
// }

package controller

import (
	"fmt"
	"log"

	"synapsis-challange/helper"
	"synapsis-challange/model/web"
	"synapsis-challange/service"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
	LogService  service.LogService
}

type AuthController interface {
	NewAuthRouter(app *fiber.App)
}

func NewAuthController(authService service.AuthService, logService service.LogService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		LogService:  logService,
	}
}

func (controller *AuthControllerImpl) NewAuthRouter(app *fiber.App) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Post("/logout", controller.Logout)
	// app.Post("/forget-password", controller.ForgetPassword)
	// app.Post("/reset-password", controller.ResetPassword)
}

func (controller *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	var request web.RegisterUserRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	_ = controller.AuthService.Register(request)

	action := fmt.Sprintf("register user %s", request.Name)
	log.Println("LOG", action)

	return ctx.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:    fiber.StatusCreated,
		Status:  true,
		Message: "Pendaftaran berhasil.",
	})
}

func (controller *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	var request web.LoginRequest
	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	cookie, data := controller.AuthService.Login(request)

	ctx.Cookie(&cookie)

	cookieData := ctx.Cookies("token")
	actor, _, _ := helper.ParseJwt(cookieData)
	action := fmt.Sprintf("login user %s", request.Email)
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func (controller *AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
	cookie := controller.AuthService.Logout()

	ctx.Cookie(&cookie)

	//produce to kafka system log
	cookieData := ctx.Cookies("token")
	actor, _, _ := helper.ParseJwt(cookieData)
	action := fmt.Sprintf("logout user %s", actor)
	controller.LogService.Create(actor, action)

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
	})
}

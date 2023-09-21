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

// func (controller *AuthControllerImpl) ForgetPassword(ctx *fiber.Ctx) error {
// 	var data web.ForgetPassowrd
// 	err := ctx.BodyParser(&data)
// 	helper.PanicIfError(err)

// 	var action string

// 	if data.Email != "" {
// 		controller.AuthService.ForgetPasswordEmail(data.Email)
// 		action = fmt.Sprintf("request reset password %s", data.Email)
// 	} else if data.Phone != "" {
// 		controller.AuthService.ForgetPasswordPhone(data.Phone)
// 		action = fmt.Sprintf("request reset password %s", data.Phone)
// 	}

// 	//produce to kafka system log
// 	cookieData := ctx.Cookies("token")
// 	actor, _, _ := helper.ParseJwt(cookieData)
// 	controller.LogService.Create(actor, action)

// 	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
// 		Code:    fiber.StatusOK,
// 		Status:  true,
// 		Message: "Permintaan reset password telah terkirim.",
// 	})
// }

// func (controller *AuthControllerImpl) ResetPassword(ctx *fiber.Ctx) error {
// 	email := ctx.Query("email")
// 	phone := ctx.Query("phone")
// 	token := ctx.Query("token")

// 	if len(token) == 0 {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Status:  false,
// 			Message: "token missing",
// 		})
// 	}
// 	if len(email) == 0 && len(phone) == 0 {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
// 			Code:    fiber.StatusBadRequest,
// 			Status:  false,
// 			Message: "query missing",
// 		})
// 	}

// 	var request web.ResetPassword
// 	err := ctx.BodyParser(&request)
// 	helper.PanicIfError(err)

// 	controller.AuthService.ResetPassword(email, phone, token, request)

// 	//produce to kafka system log
// 	cookieData := ctx.Cookies("token")
// 	actor, _, _ := helper.ParseJwt(cookieData)
// 	action := fmt.Sprintf("reset password %s", email)
// 	controller.LogService.Create(actor, action)

// 	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
// 		Code:    fiber.StatusOK,
// 		Status:  true,
// 		Message: "Reset password berhasil.",
// 	})
// }

package service

import (
	"fmt"
	"strings"
	"time"

	"synapsis-challange/exception"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
	"synapsis-challange/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thanhpk/randstr"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

type AuthService interface {
	Register(request web.RegisterUserRequest) web.UserResponse
	Login(request web.LoginRequest) (fiber.Cookie, web.LoginResponse)
	Logout() fiber.Cookie
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(request web.RegisterUserRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	if email, _ := service.AuthRepository.FindByQuery("email", request.Email); email.ID != "" {
		panic(exception.NewError(fiber.StatusBadRequest, "Email sudah terdaftar."))
	}

	if err != nil {
		panic(exception.NewError(fiber.StatusInternalServerError, "internal server error"))
	}

	user := domain.User{
		ID:        strings.ToLower(randstr.String(10)),
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.SetPassword(request.Password)

	service.AuthRepository.Register(user)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) Login(request web.LoginRequest) (fiber.Cookie, web.LoginResponse) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	var user domain.User

	if request.Email != "" {
		user, err = service.AuthRepository.LoginEmail(request.Email)
		if err != nil {
			panic(exception.NewError(fiber.StatusInternalServerError, "server error"))
		}
	}

	if user.ID == "" {
		panic(exception.NewError(fiber.StatusNotFound, "Email tidak ditemukan."))
	}

	if err := user.ComparePassword(user.Password, request.Password); err != nil {
		panic(exception.NewError(fiber.StatusBadRequest, "Password salah."))
	}

	token, err := helper.GenerateJwt(user.ID, fmt.Sprintf("%d", 1))
	if err != nil {
		helper.PanicIfError(err)
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	resp := web.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return cookie, resp
}

func (service *AuthServiceImpl) Logout() fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	return cookie
}

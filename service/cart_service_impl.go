package service

import (
	"log"
	"strings"

	"synapsis-challange/exception"
	"synapsis-challange/helper"
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
	"synapsis-challange/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thanhpk/randstr"
)

type CartServiceImpl struct {
	CartRepository     repository.CartRepository
	CategoryRepository repository.CategoryRepository
	ProductRepository  repository.ProductRepository
	Validate           *validator.Validate
}

type CartService interface {
	Create(request web.CartCreateRequest, actor string) web.CartResponse
	FindById(cartId string) web.CartResponse
	FindByUserId(userId string) web.CartResponse
	FindAll(UserID string, ProductID string) []web.CartResponse
	Update(cartId string, request web.CartUpdateRequest) web.CartResponse
	Delete(cartId string)
	DeleteByUserId(userId string)
}

func NewCartService(cartRepository repository.CartRepository, productRepository repository.ProductRepository, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		CartRepository:    cartRepository,
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (service *CartServiceImpl) Create(request web.CartCreateRequest, actor string) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// isCategoryExist := service.CategoryRepository.FindIsExist(actor, "")
	// if !isCategoryExist {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Category not exists"))
	// }

	isCartExist := service.CartRepository.FindIsExist(request.UserID, request.ProductID)
	if isCartExist {
		panic(exception.NewError(fiber.StatusBadRequest, "Product already exists in Cart"))
	}

	product, _ := service.ProductRepository.FindById(request.ProductID)
	total := request.Qty * product.Price

	cart := domain.Cart{
		ID:        strings.ToLower(randstr.String(10)),
		UserID:    request.UserID,
		ProductID: request.ProductID,
		Qty:       request.Qty,
		Total:     total,
	}

	service.CartRepository.Create(cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) FindById(cartId string) web.CartResponse {
	cart, err := service.CartRepository.FindById(cartId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Cart not found"))
	}

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) FindByUserId(cartId string) web.CartResponse {
	cart, err := service.CartRepository.FindById(cartId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Cart not found"))
	}

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) FindAll(UserID string, ProductID string) []web.CartResponse {
	carts := service.CartRepository.FindAll(UserID, ProductID)
	return helper.ToCartResponses(carts)
}

func (service *CartServiceImpl) Update(cartId string, request web.CartUpdateRequest) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// isCartExist := service.CartRepository.FindIsExist(request.Label, request.Value)
	// if isCartExist {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Cart already exists"))
	// }

	cart, err := service.CartRepository.FindById(cartId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Cart not found"))
	}

	product, _ := service.ProductRepository.FindById(cart.ProductID)
	total := request.Qty * product.Price

	// cart.UserID = request.UserID
	// cart.ProductID = request.ProductID
	cart.Qty = request.Qty
	cart.Total = total

	service.CartRepository.Update(cart, cartId)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) Delete(cartId string) {
	cart, err := service.CartRepository.FindById(cartId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Cart not found"))
	}

	err = service.CartRepository.Delete(cartId)
	helper.PanicIfError(err)
	log.Println("LOG ", "deleted cart", cart.ID)
}

func (service *CartServiceImpl) DeleteByUserId(userId string) {
	cart, err := service.CartRepository.FindByUserId(userId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Cart not found"))
	}

	err = service.CartRepository.DeleteByUserId(userId)
	helper.PanicIfError(err)
	log.Println("LOG ", "deleted cart", cart.ID)
}

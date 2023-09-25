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

type ProductServiceImpl struct {
	ProductRepository  repository.ProductRepository
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

type ProductService interface {
	Create(request web.ProductCreateRequest) web.ProductResponse
	FindById(productId string) web.ProductResponse
	FindAll(name string, category string) []web.ProductResponse
	Update(productId string, request web.ProductUpdateRequest) web.ProductResponse
	Delete(productId string)
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := domain.Product{
		ID:         strings.ToLower(randstr.String(10)),
		CategoryID: request.CategoryID,
		Name:       request.Name,
		Qty:        request.Qty,
		Price:      request.Price,
	}

	service.ProductRepository.Create(product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindById(productId string) web.ProductResponse {
	product, err := service.ProductRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Product not found"))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(name string, category string) []web.ProductResponse {
	products := service.ProductRepository.FindAll(name, category)
	return helper.ToProductResponses(products)
}

func (service *ProductServiceImpl) Update(productId string, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Product not found"))
	}

	product.CategoryID = request.CategoryID
	product.Name = request.Name
	product.Qty = request.Qty
	product.Price = request.Price

	service.ProductRepository.Update(product, productId)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(productId string) {
	product, err := service.ProductRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Product not found"))
	}

	err = service.ProductRepository.Delete(productId)
	helper.PanicIfError(err)
	log.Println("LOG ", "deleted product", product.Name)
}

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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

type CategoryService interface {
	Create(request web.CategoryCreateRequest) web.CategoryResponse
	FindById(productId string) web.CategoryResponse
	FindAll(name string) []web.CategoryResponse
	Update(productId string, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(productId string)
}

func NewCategoryService(productRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: productRepository,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := domain.Category{
		ID:   strings.ToLower(randstr.String(10)),
		Name: request.Name,
	}

	service.CategoryRepository.Create(product)

	return helper.ToCategoryResponse(product)
}

func (service *CategoryServiceImpl) FindById(productId string) web.CategoryResponse {
	product, err := service.CategoryRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Category not found"))
	}

	return helper.ToCategoryResponse(product)
}

func (service *CategoryServiceImpl) FindAll(name string) []web.CategoryResponse {
	products := service.CategoryRepository.FindAll(name)
	return helper.ToCategoryResponses(products)
}

func (service *CategoryServiceImpl) Update(productId string, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product, err := service.CategoryRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Category not found"))
	}

	product.Name = request.Name

	service.CategoryRepository.Update(product, productId)

	return helper.ToCategoryResponse(product)
}

func (service *CategoryServiceImpl) Delete(productId string) {
	product, err := service.CategoryRepository.FindById(productId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Category not found"))
	}

	err = service.CategoryRepository.Delete(productId)
	helper.PanicIfError(err)
	log.Println("LOG ", "deleted product", product.Name)
}

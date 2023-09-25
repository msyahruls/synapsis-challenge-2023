package service

import (
	"fmt"
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

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	CategoryRepository    repository.CategoryRepository
	Validate              *validator.Validate
}

type TransactionService interface {
	Create(request web.TransactionCreateRequest, actor string) web.TransactionResponse
	FindById(transactionId string) web.TransactionResponse
	FindAll(UserID string) []web.TransactionResponse
	// Update(transactionId string, request web.TransactionUpdateRequest) web.TransactionResponse
	// Delete(transactionId string)
}

func NewTransactionService(transactionRepository repository.TransactionRepository, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) Create(request web.TransactionCreateRequest, actor string) web.TransactionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// isCategoryExist := service.CategoryRepository.FindIsExist(actor, "")
	// if !isCategoryExist {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Category not exists"))
	// }

	// isTransactionExist := service.TransactionRepository.FindIsExist(request.UserID, request.ProductID)
	// if isTransactionExist {
	// 	panic(exception.NewError(fiber.StatusBadRequest, "Product already exists in Transaction"))
	// }

	// detail := []domain.TransactionDetails{
	// 	ProductID: request.ProductID,
	// 	Qty:       request.Qty,
	// 	Total:     0,
	// }

	transaction := domain.Transaction{
		ID:                 strings.ToLower(randstr.String(10)),
		UserID:             request.UserID,
		SumTotal:           request.SumTotal,
		Payment:            request.Payment,
		Payback:            request.Payback,
		TransactionDetails: request.Transactions,
	}

	fmt.Println(transaction)

	service.TransactionRepository.Create(transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindById(transactionId string) web.TransactionResponse {
	transaction, err := service.TransactionRepository.FindById(transactionId)
	if err != nil {
		panic(exception.NewError(fiber.StatusNotFound, "Transaction not found"))
	}

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindAll(UserID string) []web.TransactionResponse {
	transactions := service.TransactionRepository.FindAll(UserID)
	return helper.ToTransactionResponses(transactions)
}

// func (service *TransactionServiceImpl) Update(transactionId string, request web.TransactionUpdateRequest) web.TransactionResponse {
// 	err := service.Validate.Struct(request)
// 	helper.PanicIfError(err)

// 	// isTransactionExist := service.TransactionRepository.FindIsExist(request.Label, request.Value)
// 	// if isTransactionExist {
// 	// 	panic(exception.NewError(fiber.StatusBadRequest, "Transaction already exists"))
// 	// }

// 	transaction, err := service.TransactionRepository.FindById(transactionId)
// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusNotFound, "Transaction not found"))
// 	}

// 	// transaction.UserID = request.UserID
// 	transaction.ProductID = request.ProductID
// 	transaction.Qty = request.Qty
// 	transaction.Total = 1

// 	service.TransactionRepository.Update(transaction, transactionId)

// 	return helper.ToTransactionResponse(transaction)
// }

// func (service *TransactionServiceImpl) AssignPermission(transactionId string, request web.TransactionUpdateRequest) web.TransactionResponse {
// 	transaction, err := service.TransactionRepository.FindById(transactionId)
// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusNotFound, "Transaction not found"))
// 	}

// 	transaction.Permission = request.Permission

// 	service.TransactionRepository.AssignPermission(transaction, transactionId)

// 	return helper.ToTransactionResponse(transaction)
// }

// func (service *TransactionServiceImpl) Delete(transactionId string) {
// 	transaction, err := service.TransactionRepository.FindById(transactionId)
// 	if err != nil {
// 		panic(exception.NewError(fiber.StatusNotFound, "Transaction not found"))
// 	}

// 	err = service.TransactionRepository.Delete(transactionId)
// 	helper.PanicIfError(err)
// 	log.Println("LOG ", "deleted transaction", transaction.ID)
// }

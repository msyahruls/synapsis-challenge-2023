package service

import (
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

	transaction := domain.Transaction{
		ID:                 strings.ToLower(randstr.String(10)),
		UserID:             request.UserID,
		SumTotal:           request.SumTotal,
		Payment:            request.Payment,
		Payback:            request.Payback,
		TransactionDetails: request.Transactions,
	}

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

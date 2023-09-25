package helper

import (
	"synapsis-challange/model/domain"
	"synapsis-challange/model/web"
)

// User
func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// Category
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}

// Product
func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:         product.ID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Qty:        product.Qty,
		Price:      product.Price,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productsResponse []web.ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, ToProductResponse(product))
	}
	return productsResponse
}

// Cart
func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Qty:       cart.Qty,
		Total:     cart.Total,
	}
}

func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartsResponse []web.CartResponse
	for _, cart := range carts {
		cartsResponse = append(cartsResponse, ToCartResponse(cart))
	}
	return cartsResponse
}

// Transaction
func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	transactionDetails := ToTransactionDetailsResponses(transaction.TransactionDetails)
	return web.TransactionResponse{
		ID:                 transaction.ID,
		UserID:             transaction.UserID,
		Total:              transaction.SumTotal,
		Payment:            transaction.Payment,
		Payback:            transaction.Payback,
		TransactionDetails: transactionDetails,
	}
}

func ToTransactionResponses(transactions []domain.Transaction) []web.TransactionResponse {
	var transactionsResponse []web.TransactionResponse
	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, ToTransactionResponse(transaction))
	}
	return transactionsResponse
}

// TransactionDetails
func ToTransactionDetailsResponse(transactionDetail domain.TransactionDetails) web.TransactionDetailsResponse {
	return web.TransactionDetailsResponse{
		ProductID: transactionDetail.ProductID,
		Qty:       transactionDetail.Qty,
		Total:     transactionDetail.Total,
	}
}

func ToTransactionDetailsResponses(transactionDetails []domain.TransactionDetails) []web.TransactionDetailsResponse {
	var transactionDetailsResponse []web.TransactionDetailsResponse
	for _, transactionDetail := range transactionDetails {
		transactionDetailsResponse = append(transactionDetailsResponse, ToTransactionDetailsResponse(transactionDetail))
	}
	return transactionDetailsResponse
}

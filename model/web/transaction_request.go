package web

import "synapsis-challange/model/domain"

type TransactionCreateRequest struct {
	UserID       string                      `json:"user" validate:"required"`
	SumTotal     int                         `json:"sum_total" validate:"required"`
	Payment      int                         `json:"payment" validate:"required"`
	Payback      int                         `json:"payback" validate:"required"`
	Transactions []domain.TransactionDetails `json:"transactions" bson:"transactions"`
	// ProductID string `json:"product" validate:"required"`
	// Qty       int    `json:"qty" validate:"required"`

	// TransactionID string `json:"transaction" bson:"transaction"`
	// ProductID string `json:"product" validate:"required"`
	// Qty       string `json:"qty" validate:"required"`
	// Total     string `json:"total" validate:"required"`
}

// type TransactionUpdateRequest struct {
// 	UserID string `json:"user" validate:"required"`
// 	// 	// ProductID string `json:"product" validate:"required"`
// 	// 	// Qty       int    `json:"qty" validate:"required"`
// 	// 	// Total     string `json:"total" validate:"required"`

// 	// ProductID string `json:"product" validate:"required"`
// 	// Qty       string `json:"qty" validate:"required"`
// 	// Total     string `json:"total" validate:"required"`
// }

// type TransactionDetailsRequest struct {
// 	ProductID string `json:"product" validate:"required"`
// 	Qty       int    `json:"qty" validate:"required"`
// 	Total     int    `json:"total" validate:"required"`
// 	// TransactionID string `json:"transaction" bson:"transaction"`
// }

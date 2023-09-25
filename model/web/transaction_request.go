package web

import "synapsis-challange/model/domain"

type TransactionCreateRequest struct {
	UserID       string                      `json:"user" validate:"required"`
	SumTotal     int                         `json:"sum_total" validate:"required"`
	Payment      int                         `json:"payment" validate:"required"`
	Payback      int                         `json:"payback" validate:"required"`
	Transactions []domain.TransactionDetails `json:"transactions" bson:"transactions"`
}

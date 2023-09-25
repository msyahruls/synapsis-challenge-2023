package domain

type Transaction struct {
	ID                 string               `json:"id" bson:"_id"`
	UserID             string               `json:"user_id" bson:"user_id"`
	SumTotal           int                  `json:"sum_total" bson:"sum_total"`
	Payment            int                  `json:"payment" bson:"payment"`
	Payback            int                  `json:"payback" bson:"payback"`
	TransactionDetails []TransactionDetails `json:"transactions" bson:"transactions"`
}

type TransactionDetails struct {
	ProductID string `json:"product" bson:"product"`
	Qty       int    `json:"qty" bson:"qty"`
	Total     int    `json:"total" bson:"total"`
	// TransactionID string `json:"transaction" bson:"transaction"`
}

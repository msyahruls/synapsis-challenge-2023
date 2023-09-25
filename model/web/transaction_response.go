package web

type TransactionResponse struct {
	// ID        string `json:"id" bson:"_id"`
	// UserID    string `json:"user" bson:"user"`
	// ProductID string `json:"product_id" bson:"product_id"`
	// Qty       int    `json:"qty" bson:"qty"`
	// Total     int    `json:"total" bson:"total"`

	ID                 string                       `json:"id" bson:"_id"`
	UserID             string                       `json:"user" bson:"user"`
	Total              int                          `json:"sum_total" bson:"sum_total"`
	Payment            int                          `json:"payment" bson:"payment"`
	Payback            int                          `json:"payback" bson:"payback"`
	TransactionDetails []TransactionDetailsResponse `json:"transactions" bson:"transactions"`
}

type TransactionDetailsResponse struct {
	ProductID string `json:"product" bson:"product"`
	Qty       int    `json:"qty" bson:"qty"`
	Total     int    `json:"total" bson:"total"`
	// TransactionID string `json:"transaction" bson:"transaction"`
}

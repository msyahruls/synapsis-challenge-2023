package web

type CartResponse struct {
	ID        string `json:"id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	ProductID string `json:"product_id" bson:"product_id"`
	Qty       int    `json:"qty" bson:"qty"`
	Total     int    `json:"total" bson:"total"`
}

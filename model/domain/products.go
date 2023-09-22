package domain

type Product struct {
	ID         string `json:"id" bson:"_id"`
	CategoryID string `json:"category_id" bson:"category_id"`
	Name       string `json:"name" bson:"name"`
	Qty        int    `json:"qty" bson:"qty"`
	Price      int    `json:"price" bson:"price"`
}

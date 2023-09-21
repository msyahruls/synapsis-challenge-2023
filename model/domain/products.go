package domain

type Product struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	CategoryID string `json:"category_id" bson:"category_id"`
}

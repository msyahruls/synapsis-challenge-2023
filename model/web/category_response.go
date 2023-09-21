package web

type CategoryResponse struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

package web

type CartCreateRequest struct {
	UserID    string `json:"user" validate:"required"`
	ProductID string `json:"product" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
}

type CartUpdateRequest struct {
	UserID string `json:"user" validate:"required"`
	Qty    int    `json:"qty" validate:"required"`
}

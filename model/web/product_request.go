package web

type ProductCreateRequest struct {
	CategoryID string `json:"category" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Qty        int    `json:"qty" validate:"required"`
	Price      int    `json:"price" validate:"required"`
}

type ProductUpdateRequest struct {
	CategoryID string `json:"category" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Qty        int    `json:"qty" validate:"required"`
	Price      int    `json:"price" validate:"required"`
}

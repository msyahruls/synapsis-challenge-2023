package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" validate:"required"`
}

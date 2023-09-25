package web

type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

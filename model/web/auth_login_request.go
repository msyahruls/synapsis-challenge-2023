package web

type LoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
	// Phone    string `json:"phone"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// OccupationLevel int    `json:"occupation_level"`
	// OccupationName  string `json:"occupation_name"`
}

// type ForgetPassowrd struct {
// 	Email string `json:"email"`
// 	Phone string `json:"phone"`
// }

// type ResetPassword struct {
// 	Password        string `json:"password"`
// 	PasswordConfirm string `json:"password_confirm"`
// }

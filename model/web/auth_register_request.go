package web

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	// NIK          string `json:"nik" validate:"required"`
	// Phone        string `json:"phone" validate:"required"`
	// Gender       string `json:"gender" validate:"required"`
	// ReferralCode string `json:"referral_code"`
	// Longitude    string `json:"longitude"`
	// Latitude     string `json:"latitude"`
}

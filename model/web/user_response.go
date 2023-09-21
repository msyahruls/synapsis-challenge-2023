package web

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Gender       string `json:"gender"`
	// Phone        string `json:"phone"`
	// NIK          string `json:"nik"`
	// DistrictId   string `json:"district_id"`
	// Status       *bool  `json:"status"`
	// Longitude    string `json:"longitude"`
	// Latitude     string `json:"latitude"`
	// ProfileImage string `json:"profile_image"`
	// Occupation   OccupationResponseUser `json:"occupation"`
	// Accesses  []string  `json:"accesses"`
}

type OccupationResponseUser struct {
	Level int    `json:"level"`
	Name  string `json:"name"`
}

type UserResponseAyrshare struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	AyrshareToken string    `json:"ayrshare_token"`
	CreatedAt     time.Time `json:"created_at"`
}

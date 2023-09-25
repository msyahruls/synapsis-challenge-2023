package web

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

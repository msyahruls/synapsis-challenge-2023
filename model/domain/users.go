package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	// OccupationId string    `json:"occupation_id" bson:"occupation_id"`
	// NIK          string    `json:"nik" bson:"nik"`
	// Phone        string    `json:"phone" bson:"phone"`
	// Gender       string    `json:"gender" bson:"gender"`
	// Status       *bool     `json:"status" bson:"status"`
	// DistrictId   string    `json:"district_id" bson:"district_id"`
	// Longitude    string    `json:"longitude" bson:"longitude"`
	// Latitude     string    `json:"latitude" bson:"latitude"`
	// ReferralCode string    `json:"referral_code" bson:"referral_code"`
	// ProfileImage string    `json:"profile_image" bson:"profile_image"`
	// Occupation   Occupation  `json:"occupation" bson:"occupation"`
	// Recruitment  UserRecruit `json:"recruitment" bson:"recruitment"`
	// Roles    []string `json:"roles" bson:"roles"`
	// Accesses []string `json:"accesses" bson:"accesses"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(hashedPassword)
}

func (user *User) ComparePassword(correctPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(password))
}

type UserAyrshare struct {
	Id            string    `json:"id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`
	Email         string    `json:"email" bson:"email"`
	AyrshareToken string    `json:"ayrshare_token" bson:"ayrshare_token"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
}

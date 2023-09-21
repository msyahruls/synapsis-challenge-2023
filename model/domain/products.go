package domain

type Product struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	// OccupationId string    `json:"occupation_id" bson:"occupation_id"`
	// NIK          string    `json:"nik" bson:"nik"`
	// Email        string    `json:"email" bson:"email"`
	// Phone        string    `json:"phone" bson:"phone"`
	// Gender       string    `json:"gender" bson:"gender"`
	// Status       *bool     `json:"status" bson:"status"`
	// DistrictId   string    `json:"district_id" bson:"district_id"`
	// Longitude    string    `json:"longitude" bson:"longitude"`
	// Latitude     string    `json:"latitude" bson:"latitude"`
	// Password     string    `json:"password" bson:"password"`
	// ReferralCode string    `json:"referral_code" bson:"referral_code"`
	// ProfileImage string    `json:"profile_image" bson:"profile_image"`
	// CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	// UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	// // Occupation   Occupation  `json:"occupation" bson:"occupation"`
	// // Recruitment  UserRecruit `json:"recruitment" bson:"recruitment"`
	// Roles    []string `json:"roles" bson:"roles"`
	// Accesses []string `json:"accesses" bson:"accesses"`
}

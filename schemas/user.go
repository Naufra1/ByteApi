package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Email         string
	CriptPassword string
	Birthday      string
	Number        int64
	Address       string
	Cart          []Cart
}

type UserResponse struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     time.Time  `json:"deletedAt,omitempty"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Birthday      *time.Time `json:"birthday"`
	Number        int64      `json:"number"`
	Address       string     `json:"address"`
}

type LoginResponse struct {
	ID            uint       `json:"id"`
	Email         string     `json:"email"`
}

type ChangePasswordResponse struct {
	ID            uint       `json:"id"`
	Email         string     `json:"email"`
}

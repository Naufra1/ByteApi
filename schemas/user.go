package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Email         string
	Password      string
	CriptPassword string
	Birthday      *time.Time
	Number        int64
}

type UserResponse struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     time.Time  `json:"deletedAt,omitempty"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	CriptPassword string     `json:"criptpassword"`
	Birthday      *time.Time `json:"birthday"`
	Number        int64      `json:"number"`
}

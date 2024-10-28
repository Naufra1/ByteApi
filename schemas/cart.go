package schemas

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint `gorm:"foreignKey:UserId;references:ID"`
	// User       User
	ComputerId uint `gorm:"foreignKey:ComputerId;references:ID"`
	// Computer   Computer `gorm:"foreignKey:ComputerId;references:ID"`
}

type CartResponse struct {
	ID     uint `json:"id"`
	UserId uint `json:"userid"`
	// User       UserResponse     `json:"user"`
	ComputerId uint `json:"computerid"`
	// Computer   ComputerResponse `json:"computer"`
}

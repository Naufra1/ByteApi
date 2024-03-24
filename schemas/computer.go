package schemas

import (
	"gorm.io/gorm"
)

type Computer struct {
	gorm.Model
	Name  string
	Image string
	Price string
	Specs map[string]string
}

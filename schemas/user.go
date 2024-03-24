package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	CriptPasword string
	Birthday     *time.Time
	Number       int64
}

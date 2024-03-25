package handler

import (
	"github.com/Naufra1/ByteApi/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

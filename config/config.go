package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Inicializando o banco de dados
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("erro ao inicializar o SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

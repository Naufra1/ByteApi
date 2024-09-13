package config

import (
	"os"

	"github.com/Naufra1/ByteApi/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Checando se o banco de dados existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Arquivo de banco de dados não encontrado, criando...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Criando e conectando o banco de dados
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro na inicialização do banco de dados: %v", err)
		return nil, err
	}

	// Migrando os schemas
	err = db.AutoMigrate(&schemas.Computer{})
	if err != nil {
		logger.Errorf("Erro na migração de computadores: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("Erro na migração de usuários: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Cart{})
	if err != nil {
		logger.Errorf("Erro na migração de carrinhos: %v", err)
	}

	return db, nil
}

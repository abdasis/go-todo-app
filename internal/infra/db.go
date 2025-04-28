package infra

import (
	"go-todo-app/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DatabaseName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return db
}

package infra

import (
	"go-todo-app/internal/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
}

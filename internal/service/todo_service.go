package service

import (
	"go-todo-app/internal/model"
	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) CreateTodo(todo *model.Todo) error {
	return s.db.Create(todo).Error
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	err := s.db.Find(&todos).Error
	return todos, err
}

func (s *TodoService) GetTodoByID(id uint) (model.Todo, error) {
	var todo model.Todo
	err := s.db.First(&todo, id).Error
	return todo, err
}

func (s *TodoService) UpdateTodo(todo *model.Todo) error {
	return s.db.Save(todo).Error
}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.db.Delete(&model.Todo{}, id).Error
}

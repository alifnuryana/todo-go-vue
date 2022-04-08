package repository

import (
	"errors"

	"github.com/alifnuryana/go-auth-jwt/models"
	"gorm.io/gorm"
)

type TodoInterface interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id int) (models.Todo, error)
	CreateTodo(todo models.Todo) error
	UpdateTodo(todo models.Todo, id int) error
	DeleteTodo(id int) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoInterface {
	return &TodoRepository{db: db}
}

func (t *TodoRepository) GetTodos() ([]models.Todo, error) {
	var todo []models.Todo
	tx := t.db.Find(&todo)
	if tx.RowsAffected == 0 {
		return todo, errors.New("error record not found")
	}
	return todo, nil
}

func (t *TodoRepository) GetTodo(id int) (models.Todo, error) {
	var todo models.Todo
	tx := t.db.First(&todo, id)
	if tx.RowsAffected == 0 {
		return todo, errors.New("error record not found")
	}
	return todo, nil
}

func (t *TodoRepository) CreateTodo(todo models.Todo) error {
	tx := t.db.Create(&todo)
	return tx.Error
}

func (t *TodoRepository) UpdateTodo(todo models.Todo, id int) error {
	tx := t.db.Model(models.Todo{}).Where("id = ?", id).Updates(todo)
	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return tx.Error
}

func (t *TodoRepository) DeleteTodo(id int) error {
	tx := t.db.Delete(&models.Todo{}, id)
	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return tx.Error
}

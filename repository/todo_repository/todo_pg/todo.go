package todo_pg

import (
	"gorm.io/gorm"
	"todo/entity"
	"todo/pkg/errs"
	"todo/repository/todo_repository"
)

type todoPg struct {
	db *gorm.DB
}

func NewTodoPG(database *gorm.DB) todo_repository.TodoRepository {
	return &todoPg{db: database}
}

func (t *todoPg) CreateTodo(todoPayload *entity.Todo) (*entity.Todo, errs.MessageErr) {
	if err := t.db.Create(todoPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to create data")
	}
	return todoPayload, nil
}

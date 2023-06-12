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

func (t *todoPg) GetAllTodo() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo

	if err := t.db.Find(&todos).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to get data")
	}

	return todos, nil
}

func (t *todoPg) GetTodoById(id uint) (*entity.Todo, errs.MessageErr) {
	var todo entity.Todo
	if err := t.db.First(&todo, id).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find data")
	}

	return &todo, nil
}

func (t *todoPg) UpdateTodoById(id uint, todoPayload *entity.Todo) (*entity.Todo, errs.MessageErr) {
	todo, err := t.GetTodoById(id)
	if err != nil {
		return nil, errs.NewNotFound("Error occurred while trying to find data")
	}

	if err := t.db.Model(todo).Updates(todoPayload).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying update data")
	}

	return todo, nil
}

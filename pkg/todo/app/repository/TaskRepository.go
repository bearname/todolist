package repository

import "todolist/pkg/todo/app/model"

type TaskRepository interface {
	InsertTodo(description string) error
	MarkTaskAsCompleted(id string, isCompleted bool) error
	GetNotCompletedTask() ([]model.Task, error)
	GetCompletedTask() ([]model.Task, error)
	DeleteTask(id string) error
}
package service

import "todolist/pkg/todo/app/model"

type TaskService interface {
	InsertTodo(description string) error
	MarkTaskAsCompleted(id string) error
	GetNotCompletedTask() ([]model.Task, error)
	DeleteTask(id string) bool
}

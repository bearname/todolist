package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todolist/pkg/todo/app/model"
	"todolist/pkg/todo/app/repository"
)

type TaskController struct {
	repository repository.TaskRepository
}

func NewTaskController(repository repository.TaskRepository) *TaskController {
	controller := new(TaskController)
	controller.repository = repository
	return controller
}

func (c *TaskController) CreateBook(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t model.TaskDto
	err := decoder.Decode(&t)
	if err != nil {
		Error(writer, err, http.StatusInternalServerError)
		return
	}
	log.Println(t.Description + " ")
	if len(t.Description) > 1000 {
		Error(writer, err, http.StatusNotFound)
		return
	}
	book := c.repository.InsertTodo(t.Description)
	JsonResponse(writer, book)
}

func (c *TaskController) MarkTaskAsCompleted(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	err := c.repository.MarkTaskAsCompleted(id, true)

	var result bool
	if err != nil {
		result = false
	} else {
		result = true
	}

	JsonResponse(writer, result)
}

func (c *TaskController) GetNotCompletedTask(writer http.ResponseWriter, _ *http.Request) {
	c.completedTask(writer, false)
}

func (c *TaskController) GetCompletedTask(writer http.ResponseWriter, _ *http.Request) {
	c.completedTask(writer, true)
}

func (c *TaskController) completedTask(writer http.ResponseWriter, isCompleted bool) {
	var task []model.Task
	var err error

	if isCompleted {
		task, err = c.repository.GetCompletedTask()
	} else {
		task, err = c.repository.GetNotCompletedTask()

	}
	if err != nil {
		Error(writer, err, http.StatusInternalServerError)
		return
	}
	jsonTasks, err := json.Marshal(task)
	if err != nil {
		Error(writer, err, http.StatusBadRequest)
		return
	}

	JsonResponse(writer, string(jsonTasks))
}

func (c *TaskController) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	err := c.repository.DeleteTask(id)

	var result bool
	if err != nil {
		result = false
	} else {
		result = true
	}

	JsonResponse(writer, result)
}

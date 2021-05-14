package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todolist/pkg/todo/app/model"
	"todolist/pkg/todo/app/repository"
	"todolist/pkg/todo/infrastructure/util"
)

type TaskController struct {
	repository repository.TaskRepository
	BaseController
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
		c.BaseController.Error(writer, err, http.StatusInternalServerError)
		return
	}
	log.Println(t.Description + " ")
	if len(t.Description) > 1000 {
		c.BaseController.Error(writer, err, http.StatusNotFound)
		return
	}
	book := c.repository.InsertTodo(t.Description)
	c.BaseController.JsonResponse(writer, book)
}

func (c *TaskController) MarkTaskAsCompleted(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	if c.validateId(writer, id) {
		return
	}
	err := c.repository.MarkTaskAsCompleted(id, true)

	var result bool
	if err != nil {
		result = false
	} else {
		result = true
	}

	c.BaseController.JsonResponse(writer, result)
}

func (c *TaskController) validateId(writer http.ResponseWriter, id string) bool {
	if !util.IsValidUUID(id) {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid id"))
		return true
	}
	return false
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
		c.BaseController.Error(writer, err, http.StatusInternalServerError)
		return
	}
	jsonTasks, err := json.Marshal(task)
	if err != nil {
		c.BaseController.Error(writer, err, http.StatusBadRequest)
		return
	}

	c.BaseController.JsonResponse(writer, string(jsonTasks))
}

func (c *TaskController) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if c.validateId(writer, id) {
		return
	}
	err := c.repository.DeleteTask(id)

	var result bool
	if err != nil {
		result = false
	} else {
		result = true
	}

	c.BaseController.JsonResponse(writer, result)
}

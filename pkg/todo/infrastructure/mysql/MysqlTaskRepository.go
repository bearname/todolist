package mysql

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"todolist/pkg/todo/app/model"
)

type TaskRepository struct {
	connector Connector
}

func NewTaskRepository(connector Connector) *TaskRepository {
	repository := new(TaskRepository)
	repository.connector = connector
	return repository
}

func (m TaskRepository) InsertTodo(description string) error {
	query, err := m.connector.Database.Query("INSERT INTO task (description) VALUES (?)", description)
	if err != nil {
		return err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Error(err)
		}
	}(query)
	return nil
}

func (m TaskRepository) MarkTaskAsCompleted(id string, isCompleted bool) error {
	query, err := m.connector.Database.Query("UPDATE task SET status = ?, completed_date = NOW()  WHERE id_task = ?;", isCompleted, id)
	if err != nil {
		return err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Error(err)
		}
	}(query)
	return nil
}

func (m TaskRepository) GetNotCompletedTask() ([]model.Task, error) {
	return m.getCompleted(false)
}

func (m TaskRepository) GetCompletedTask() ([]model.Task, error) {
	return m.getCompleted(true)
}

func (m TaskRepository) getCompleted(isCompleted bool) ([]model.Task, error) {
	rows, err := m.connector.Database.Query("SELECT id_task, description, status, created_date FROM task WHERE status = ?;", isCompleted)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error(err)
		}
	}(rows)

	tasks := make([]model.Task, 0)
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.Id,
			&task.Description,
			&task.Status,
			&task.CreatedDate,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (m TaskRepository) DeleteTask(id string) error {
	query, err := m.connector.Database.Query("DELETE FROM task WHERE id_task = ?;", id)
	if err != nil {
		return err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Error(err)
		}
	}(query)
	return nil
}

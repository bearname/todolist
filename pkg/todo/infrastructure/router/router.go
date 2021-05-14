package router

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todolist/pkg/todo/infrastructure/controller"
	"todolist/pkg/todo/infrastructure/mysql"
)

func Router(connector mysql.Connector) http.Handler {
	router := mux.NewRouter()
	repository := mysql.NewTaskRepository(connector)
	taskController := controller.NewTaskController(repository)
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	subrouter.HandleFunc("/task", taskController.CreateBook).Methods(http.MethodPost)
	subrouter.HandleFunc("/task/not-completed", taskController.GetNotCompletedTask).Methods(http.MethodGet)
	subrouter.HandleFunc("/task/completed", taskController.GetCompletedTask).Methods(http.MethodGet)
	subrouter.HandleFunc("/task/{id}", taskController.MarkTaskAsCompleted).Methods(http.MethodPost)
	subrouter.HandleFunc("/task/{id}", taskController.DeleteTask).Methods(http.MethodDelete)

	return logMiddleware(router)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

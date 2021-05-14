package infrastructure

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todolist/pkg/todo/infrastructure/mysql"
	"todolist/pkg/todo/infrastructure/router"
)

type Server struct {
}

func (s *Server) StartServer(serverUrl string, connector mysql.Connector) *http.Server {
	router := router.Router(connector)
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Error(srv.ListenAndServe())
	}()

	return srv
}

func (s *Server) GetKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func (s *Server) WaitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

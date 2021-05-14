package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
	"todolist/pkg/todo/infrastructure"
	"todolist/pkg/todo/infrastructure/mysql"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("todolist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	config, err := parseConfig()
	if err != nil {
		log.Info("Default settings" + err.Error())
	}
	var connector mysql.Connector
	if err == nil {
		log.Info("config " + config.DbName + config.DbUser + config.DbPassword + config.DbAddress + config.DbMigrationsDir)

		err = connector.Connect(config.DbUser, config.DbPassword, config.DbAddress, config.DbName)
		log.Info("*mysql.NewConnector")
		if err != nil {
			log.Error("unable to connect to database" + err.Error())
			return
		}
		defer connector.Close()
	} else {
		log.Info("else NewConnector")
		config.DbMigrationsDir = ""
		connector = *mysql.NewConnector("root", "123", "127.0.0.1", "todo")
		err = connector.Connect("root", "123", "127.0.0.1", "todo")
		if err != nil {
			log.Error("unable to connect to database" + err.Error())
			return
		}
		defer connector.Close()
	}

	server := infrastructure.Server{}
	killSignalChan := server.GetKillSignalChan()

	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	err = mysql.Migrate(connector, config.DbMigrationsDir)
	if err != nil {
		log.Error(err)
	}
	srv := server.StartServer(serverUrl, connector)

	server.WaitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

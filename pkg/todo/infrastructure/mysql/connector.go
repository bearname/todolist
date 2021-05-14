package mysql

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type Connector struct {
	Database *sql.DB
}

func NewConnector(dbUser string, dbPassword string, dbAddress string, dbName string) *Connector {
	var connector Connector
	err := connector.Connect(dbUser, dbPassword, dbAddress, dbName)
	if err != nil {
		panic("unable to connect to database" + err.Error())
	}

	defer connector.Close()
	return &connector
}

func (c *Connector) Connect(user string, password string, address string, databaseName string) error {
	if c.Database != nil {
		log.Info("Already connected")
	}

	connection := user + ":" + password + "@tcp(" + address + ")/" + databaseName
	log.Info("connection " + connection)
	database, err := sql.Open("mysql", connection)
	if err != nil {
		log.Error(err)
		return err
	}

	c.Database = database

	return nil
}

func (c *Connector) Close() error {
	err := c.Database.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	c.Database = nil

	return nil
}

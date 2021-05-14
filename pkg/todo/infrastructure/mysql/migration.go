package mysql

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func Migrate(connector Connector, directory string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Error(err)
		return err
	}

	var buffer strings.Builder
	for _, f := range files {
		fmt.Println(f.Name())
		bytes, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Error(err)
		}
		buffer.Write(bytes)
	}

	log.Info(buffer.String())
	query, err := connector.Database.Query(buffer.String())
	if err != nil {
		log.Error(err)
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Error(err)
		}
	}(query)
	return nil
}

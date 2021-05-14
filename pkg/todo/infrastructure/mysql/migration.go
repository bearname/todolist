package mysql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func Migrate(connector Connector, directory string)  error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
		bytes, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Error(err)
		}
		queryString := string(bytes)
		log.Info(queryString)
		query, err := connector.Database.Query(queryString)
		if err != nil {
			log.Error(err)
		}
		defer query.Close()
	}

	return nil
}

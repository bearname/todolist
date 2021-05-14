package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type config struct {
	DbName          string
	DbAddress       string
	DbUser          string
	DbPassword      string
	DbMigrationsDir string
}

func parseEnvString(key string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	str, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("undefined environment variable %v", key))
	}
	return str, nil
}

func parseConfig() (*config, error) {
	var err error
	dbName, err := parseEnvString("DATABASE_NAME", err)
	dbAddress, err := parseEnvString("DATABASE_ADDRESS", err)
	dbUser, err := parseEnvString("DATABASE_USER", err)
	dbPassword, err := parseEnvString("DATABASE_PASSWORD", err)
	dbMigrationsDir, err := parseEnvString("DATABASE_MIGRATIONS_DIR", err)

	if err != nil {
		log.Info("erro" + err.Error())
		return nil, err
	}

	return &config{
		dbName,
		dbAddress,
		dbUser,
		dbPassword,
		dbMigrationsDir,
	}, nil
}

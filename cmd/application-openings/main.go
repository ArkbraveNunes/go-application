package main

import (
	"go-basic/internal/application-openings/application/router"
	"go-basic/internal/application-openings/config/env"
	"go-basic/internal/application-openings/infra/schema"
	"go-basic/pkg/database/mysql"
	"go-basic/pkg/logger"
)

func main() {
	loggerInstance := logger.GetLogger("main")

	appEnv := env.GetAppEnv()
	databaseEnv := env.GetDatabaseEnv()

	mysqlConfig := mysql.MySqlConfig{
		Host:         databaseEnv.Host,
		Port:         databaseEnv.Port,
		DatabaseName: databaseEnv.Name,
		Username:     databaseEnv.Username,
		Password:     databaseEnv.Password,
	}

	err := mysql.MySqlInit(mysqlConfig)

	if err != nil {
		loggerInstance.ErrorWithValues("main - mysql error: %v", err)
		return
	}

	err = mysql.MySqlAutoMigrate(schema.Opening{})

	if err != nil {
		loggerInstance.ErrorWithValues("main - mysql auto migrate error: %v", err)
		return
	}

	router.InitRouter(appEnv.Port)
}

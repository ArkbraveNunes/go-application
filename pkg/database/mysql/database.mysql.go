package mysql

import (
	"fmt"
	logger "go-basic/pkg/logger"

	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

type MySqlConfig struct {
	Host         string
	Port         string
	DatabaseName string
	Username     string
	Password     string
}

var (
	database       *gorm.DB
	loggerInstance *logger.Logger
)

func GetDatabase() *gorm.DB {
	return database
}

func MySqlInit(config MySqlConfig) error {
	var err error

	database, err = mySqlInit(config)
	//database, err = mySqlInit("localhost", "3306", "application-openings", "local", "local")
	if err != nil {
		return fmt.Errorf("DatabaseMySqlInit - error: %v", err)
	}

	return nil
}

func MySqlAutoMigrate(schema interface{}) error {
	err := database.AutoMigrate(&schema)
	if err != nil {
		loggerInstance.ErrorWithValues("DatabaseAutoMigrate - auto migrate error: %v", err)
		return fmt.Errorf("DatabaseAutoMigrate - error: %v", err)
	}

	return nil
}

func mySqlInit(config MySqlConfig) (*gorm.DB, error) {
	logger := logger.GetLogger("Init Mysql Database")

	databaseString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)

	database, err := gorm.Open(mysql.Open(databaseString), &gorm.Config{})

	if err != nil {
		logger.ErrorWithValues("mySqlInit - gorm open error: %v", err)
		return nil, err
	}

	return database, nil
}

package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitMySQL()
	if err != nil {
		return fmt.Errorf("Init Mysql database error: %v", err)
	}

	return nil
}

func GetMySql() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

package config

import (
	"application-openings/schema"

	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

func InitMySQL() (*gorm.DB, error) {
	logger := GetLogger("mySQL")

	db, err := gorm.Open(mysql.Open("local:local@tcp(localhost:3306)/application-openings?parseTime=true"), &gorm.Config{})

	if err != nil {
		logger.ErrorWithValues("mysql database opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.ErrorWithValues("mysql auto migrate error: %v", err)
	}

	return db, nil
}

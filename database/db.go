package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := mysql.Open("root:@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local")
	DB, err = gorm.Open(dsn, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
}

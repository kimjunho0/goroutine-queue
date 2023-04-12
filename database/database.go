package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"queue/models"
)

const (
	User   = "root"
	Pw     = ""
	Local  = "127.0.0.1:3306"
	DBName = "queue"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local",
		User,
		Pw,
		Local,
		DBName,
	)

	dbConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	DB, err = gorm.Open(mysql.Open(dsn), dbConfig)
	if err != nil {
		panic(err)
	}
	createDB()

}

func createDB() {
	tables := []interface{}{
		(*models.Queue)(nil),
	}
	if err := DB.AutoMigrate(tables...); err != nil {
		panic(err)
	}
}

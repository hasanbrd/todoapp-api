package database

import (
	"fmt"
	"log"
	"todo-list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MYSQL_HOST     = "localhost"
	MYSQL_PORT     = "3306"
	MYSQL_USER     = "root"
	MYSQL_PASSWORD = "hasan123"
	MYSQL_DATABASE = "todoapp"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to db:", err)
	}

	db.Debug().AutoMigrate(models.ToDo{}, models.Activity{})
}

func GetDB() *gorm.DB {
	return db
}

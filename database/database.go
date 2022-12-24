package database

import (
	"fmt"
	"log"
	"os"
	"todo-list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	if MYSQL_PORT == "" {
		MYSQL_PORT = "3306"
	}

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), MYSQL_PORT, os.Getenv("MYSQL_DBNAME"))

	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to db:", err)
	}

	db.Debug().AutoMigrate(models.ToDo{}, models.Activity{})
}

func GetDB() *gorm.DB {
	return db
}

package database

import (
	"blog/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {

	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB connected")
	err = db.AutoMigrate(&models.Users{}, &models.Posts{})
	if err != nil {
		panic(err)
	}

	DBConn = db
}

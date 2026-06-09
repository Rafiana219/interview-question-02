package config

import (
	"fmt"
	"os"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Printf("host=%s port=%s user=%s db=%s\n", host, port, user, dbname)

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	fmt.Println("DB connect")

	db.AutoMigrate(&models.User{})

	DB = db
}

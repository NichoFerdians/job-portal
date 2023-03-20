package db

import (
	"fmt"
	"log"
	"os"

	"github.com/NichoFerdians/job-portal/helpers"
	"github.com/NichoFerdians/job-portal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Con *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_SERVICE")
	PORT := os.Getenv("DB_PORT")
	NAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, NAME)

	fmt.Println(URL)
	database, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	database.AutoMigrate(&models.User{})

	hashedPassword, err := helpers.HashPassword("123456")
	if err != nil {
		log.Fatal(err.Error())
	}

	newUser := &models.User{
		UserName: "admin",
		Password: hashedPassword,
		Name:     "Admin",
	}
	database.Create(newUser)

	Con = database
}

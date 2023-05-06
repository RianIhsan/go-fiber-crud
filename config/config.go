package config

import (
	"fmt"
	"log"
	"os"

	"github.com/RianIhsan/go-rest-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Tidak bisa memuat file .env")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("tidak bisa terkoneksi ke database")
	} else {
		fmt.Println("Berhasil terkoneksi ke database")
	}

	database.AutoMigrate(&models.User{})

	DB = database

}

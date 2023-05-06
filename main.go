package main

import (
	"log"
	"os"

	"github.com/RianIhsan/go-rest-api/config"
	"github.com/RianIhsan/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//Koneksi ke database
	config.ConnectDB()

	//Load .ENV file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Tidak bisa memuat file .env")
	}
	port := os.Getenv("PORT")

	//Fiber Init
	app := fiber.New()

	//Route Init
	routes.RouteInit(app)

	app.Listen(":" + port)
}

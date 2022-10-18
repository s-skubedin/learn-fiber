package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/s-skubedin/todo-fiber/models"
	"github.com/s-skubedin/todo-fiber/routes"
	"github.com/s-skubedin/todo-fiber/storage"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading file dotenv. Error: ", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error init DB. Error: ", err)
	}

	err = models.MigrateBooks(db)

	if err != nil {
		log.Fatal("Could not migrate")
	}

	r := routes.Repository{
		DB: db,
	}

	app := fiber.New()

	r.SetupBooksRoutes(app)

	app.Listen(":3000")
}

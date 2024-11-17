package main

import (
	"log"
	"minio-learn/handlers"
	"minio-learn/lib"
	"minio-learn/repository"
	"minio-learn/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := lib.ConnMinio()

	if err != nil {
		panic(err)
	}

	app := fiber.New()
	routes := app.Group("/api/v1")
	repository := repository.NewUploadRepository(client)
	serv := service.NewUploadService(repository)
	handlers := handlers.NewUploadHandler(serv)

	routes.Post("/upload", handlers.UploadFile)
	routes.Get("/file", handlers.GetFile)
	log.Fatalln(app.Listen(":8081"))

}

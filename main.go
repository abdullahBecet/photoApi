package main

import (
	"photoApi/app"
	"photoApi/configs"
	"photoApi/repository"
	"photoApi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	ClientDB := configs.GetCollection(configs.DB, "photo")
	PhotoRepositoryDB := repository.NewPhotoRepositoryDb(ClientDB)

	td := app.PhotoHandler{Service: services.NewPhotoService(PhotoRepositoryDB)}

	appRoute.Get("/api/photos", td.GetAllPhoto)

	appRoute.Listen(":8080")
}

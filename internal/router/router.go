package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wawancallahan/go-upload/internal/controller"
	"github.com/wawancallahan/go-upload/internal/service"
)

func New() *fiber.App {
	api := fiber.New()

	route := api.Group("/upload")

	uploadService := service.NewUploadService()
	uploadController := controller.NewUploadController(uploadService)

	route.Post("/", uploadController.Upload)

	return api
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wawancallahan/go-upload/internal/service"
)

type UploadController interface {
	Upload(ctx *fiber.Ctx) error
}

type UploadControllerImpl struct {
	UploadService service.UploadService
}

func NewUploadController(UploadService service.UploadService) UploadController {
	return &UploadControllerImpl{
		UploadService: UploadService,
	}
}

func (c *UploadControllerImpl) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if file == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "File Required",
		})
	}

	_, err = c.UploadService.Upload(file)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	_ = c.UploadService.Remove(file)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"error":  nil,
	})
}

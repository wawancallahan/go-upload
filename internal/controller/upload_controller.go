package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wawancallahan/go-upload/internal/service"
)

type UploadControllerImpl struct {
	UploadServiceImpl *service.UploadServiceImpl
}

func NewUploadController(UploadServiceImpl *service.UploadServiceImpl) *UploadControllerImpl {
	return &UploadControllerImpl{
		UploadServiceImpl: UploadServiceImpl,
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

	_, err = c.UploadServiceImpl.Upload(file)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	_ = c.UploadServiceImpl.Remove(file)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"error":  nil,
	})
}

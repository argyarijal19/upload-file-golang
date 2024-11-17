package handlers

import (
	"minio-learn/model"
	"minio-learn/service"

	"github.com/gofiber/fiber/v2"
)

type uploadHandlers struct {
	serv service.UploadService
}

func NewUploadHandler(serv service.UploadService) *uploadHandlers {
	return &uploadHandlers{serv}
}

func (u *uploadHandlers) UploadFile(ctx *fiber.Ctx) error {

	form, err := ctx.MultipartForm()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid form-data: " + err.Error(),
		})
	}

	fileHeaders := form.File["attachment"]
	if len(fileHeaders) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File not found in form-data",
		})
	}

	fileHeader := fileHeaders[0]
	fileName := ctx.FormValue("file_name")
	if fileName == "" {
		fileName = fileHeader.Filename
	}

	// Create the UploadFile model
	request := model.UploadFile{
		FileName: fileName,
		File:     fileHeader,
	}

	err = u.serv.UploadFile(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
	})

}

func (u *uploadHandlers) GetFile(ctx *fiber.Ctx) error {
	var request model.Request

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fileUrl, err := u.serv.GetFile(request.FileName)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File retrieved successfully",
		"url":     fileUrl,
	})

	return res
}

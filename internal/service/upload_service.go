package service

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/go-resty/resty/v2"
)

type UploadService interface {
	Upload(file *multipart.FileHeader) (bool, error)
	Remove(file *multipart.FileHeader) error
}

type UploadServiceImpl struct{}

func NewUploadService() UploadService {
	return &UploadServiceImpl{}
}

func (c *UploadServiceImpl) Upload(file *multipart.FileHeader) (bool, error) {
	buffer, err := file.Open()

	if err != nil {
		return false, err
	}

	defer buffer.Close()

	client := resty.New()

	response, err := client.R().
		SetFileReader("file", file.Filename, buffer).
		SetFormData(map[string]string{
			"filename": file.Filename,
			"folder":   "test",
			"isPublic": "true",
		}).
		Post("storage/objects/upload")

	if err != nil {
		return false, err
	}

	if response.StatusCode() > 300 {
		return false, errors.New("Response Status Over 300")
	}

	body := response.Body()

	log.Println(string(body))

	return true, nil
}

func (c *UploadServiceImpl) Remove(file *multipart.FileHeader) error {
	client := resty.New()

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"fullPathObject": fmt.Sprintf("test/%s", file.Filename),
		}).
		Delete("storage/objects/delete")

	if err != nil {
		return err
	}

	if response.StatusCode() > 300 {
		return errors.New("Response Status Over 300")
	}

	body := response.Body()

	log.Println(string(body))

	return nil
}

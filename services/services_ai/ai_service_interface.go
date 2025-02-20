package services_ai

import (
	"mime/multipart"
)

type AiServiceInterface interface {
	MakeQuestionFromFile(fileURI string, amount int) (string, error)
	StoreFileuploadFile(file multipart.File, fileName string) (string, error)
}

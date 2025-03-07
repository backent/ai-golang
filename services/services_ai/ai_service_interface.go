package services_ai

import (
	"mime/multipart"
)

type AiServiceInterface interface {
	MakeQuestionFromFile(fileURI string, amount int, chapter string, language string) (string, error)
	CheckFileMaterialExists(fileURI string, text string) (string, error)
	StoreFileuploadFile(file multipart.File, fileName string) (string, error)
}

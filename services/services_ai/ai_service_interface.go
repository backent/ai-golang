package services_ai

import (
	"mime/multipart"

	"github.com/backent/ai-golang/web/web_question"
)

type AiServiceInterface interface {
	MakeQuestionFromFile(fileURI string, description string) web_question.Result
	StoreFileuploadFile(file multipart.File, fileName string) (string, error)
}

package services_question

import (
	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/repositories/repositories_storage"
	"github.com/backent/ai-golang/services/services_ai"
	"github.com/backent/ai-golang/web/web_question"
)

type QuestionServiceImplementation struct {
	services_ai.AiServiceInterface
	repositories_storage.StorageServiceInterface
}

func NewQuestionServiceImplementation(ai services_ai.AiServiceInterface, storage repositories_storage.StorageServiceInterface) QuestionServiceInterface {
	return &QuestionServiceImplementation{
		AiServiceInterface:      ai,
		StorageServiceInterface: storage,
	}
}

func (implementation *QuestionServiceImplementation) Create(request web_question.QuestionPostRequest) web_question.Result {
	err := implementation.StorageServiceInterface.SaveFile(request.File, request.FileHeader.Filename, "storage/pdf")
	helpers.PanicIfError(err)
	fileURI, err := implementation.AiServiceInterface.StoreFileuploadFile(request.File, request.FileHeader.Filename)
	helpers.PanicIfError(err)

	data := implementation.AiServiceInterface.MakeQuestionFromFile(fileURI, request.Description)

	return data
}

package services_question

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/models"
	"github.com/backent/ai-golang/repositories/repositories_question"
	"github.com/backent/ai-golang/repositories/repositories_storage"
	"github.com/backent/ai-golang/services/services_ai"
	"github.com/backent/ai-golang/web/web_question"
)

type QuestionServiceImplementation struct {
	services_ai.AiServiceInterface
	repositories_storage.StorageServiceInterface
	repositories_question.RepositoryQuestionInterface
	*sql.DB
}

func NewQuestionServiceImplementation(ai services_ai.AiServiceInterface, storage repositories_storage.StorageServiceInterface, question repositories_question.RepositoryQuestionInterface, sql *sql.DB) QuestionServiceInterface {
	return &QuestionServiceImplementation{
		AiServiceInterface:          ai,
		StorageServiceInterface:     storage,
		RepositoryQuestionInterface: question,
		DB:                          sql,
	}
}

func (implementation *QuestionServiceImplementation) Create(ctx context.Context, request web_question.QuestionPostRequest) web_question.QuestionGetAllRequestItem {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = implementation.StorageServiceInterface.SaveFile(request.File, request.FileHeader.Filename, "storage/pdf")
	helpers.PanicIfError(err)
	fileURI, err := implementation.AiServiceInterface.StoreFileuploadFile(request.File, request.FileHeader.Filename)
	helpers.PanicIfError(err)

	textResponse, err := implementation.AiServiceInterface.MakeQuestionFromFile(fileURI, request.Amount, request.Chapter, request.Language)
	helpers.PanicIfError(err)
	var username string
	username, ok := ctx.Value(helpers.ContextKey("username")).(string)
	if !ok {
		panic("wrong username type")
	}

	questionModel := models.Question{
		Username:      username,
		Chapter:       sql.NullString{String: request.Chapter, Valid: request.Chapter != ""},
		Language:      sql.NullString{String: request.Language, Valid: request.Language != ""},
		Name:          request.Name,
		Amount:        request.Amount,
		GeminiFileURI: fileURI,
		Result:        textResponse,
		FileName:      request.FileHeader.Filename,
	}

	question, err := implementation.RepositoryQuestionInterface.Create(ctx, tx, questionModel)
	helpers.PanicIfError(err)

	return web_question.QuestionModelToQuestionGetAllRequestItem(question)
}

func (implementation *QuestionServiceImplementation) GetAll(ctx context.Context) web_question.QuestionGetAllRequest {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	collections, err := implementation.RepositoryQuestionInterface.GetAll(ctx, tx)
	helpers.PanicIfError(err)

	return web_question.CollectionQuestionModelToQuestionGetAllRequest(collections)
}

func (implementation *QuestionServiceImplementation) GetById(ctx context.Context, id int) web_question.QuestionGetByIdResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	model, err := implementation.RepositoryQuestionInterface.GetByIdWithExams(ctx, tx, id)
	helpers.PanicIfError(err)

	responseData, err := web_question.QuestionModelToQuestionGetByIdResponse(model)
	helpers.PanicIfError(err)

	return responseData
}

func (implementation *QuestionServiceImplementation) DeleteById(ctx context.Context, id int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = implementation.RepositoryQuestionInterface.DeleteById(ctx, tx, id)
	helpers.PanicIfError(err)
}

func (implementation *QuestionServiceImplementation) CheckMaterial(ctx context.Context, request web_question.QuestionCheckFileMaterialRequest) web_question.QuestionCheckMaterialResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	err = implementation.StorageServiceInterface.SaveFile(request.File, request.FileHeader.Filename, "storage/pdf")
	helpers.PanicIfError(err)
	fileURI, err := implementation.AiServiceInterface.StoreFileuploadFile(request.File, request.FileHeader.Filename)
	helpers.PanicIfError(err)

	textResponseMaterial, err := implementation.AiServiceInterface.CheckFileMaterialExists(fileURI, request.Chapter)
	helpers.PanicIfError(err)

	var checkMaterialResponse web_question.QuestionCheckMaterialResponse

	err = json.Unmarshal([]byte(textResponseMaterial), &checkMaterialResponse)
	helpers.PanicIfError(err)

	return checkMaterialResponse
}

package web_question

import (
	"encoding/json"

	"github.com/backent/ai-golang/models"
)

type QuestionGetAllRequestItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type QuestionGetAllRequest []QuestionGetAllRequestItem

func QuestionModelToQuestionGetAllRequestItem(model models.Question) QuestionGetAllRequestItem {
	return QuestionGetAllRequestItem{
		Id:   model.Id,
		Name: model.Name,
	}
}

func CollectionQuestionModelToQuestionGetAllRequest(collection []models.Question) QuestionGetAllRequest {
	var collectionRequest QuestionGetAllRequest
	for _, item := range collection {
		collectionRequest = append(collectionRequest, QuestionModelToQuestionGetAllRequestItem(item))
	}

	return collectionRequest
}

type QuestionGetByIdResponse struct {
	Id       int64        `json:"id"`
	Name     string       `json:"name"`
	Amount   int          `json:"amount"`
	FileName string       `json:"file_name"`
	Result   []ItemResult `json:"result"`
}

func QuestionModelToQuestionGetByIdResponse(model models.Question) (QuestionGetByIdResponse, error) {
	var result Result
	err := json.Unmarshal([]byte(model.Result), &result)
	if err != nil {
		return QuestionGetByIdResponse{}, err
	}
	return QuestionGetByIdResponse{
		Id:       model.Id,
		Name:     model.Name,
		Amount:   model.Amount,
		FileName: model.FileName,
		Result:   result.Result,
	}, nil
}

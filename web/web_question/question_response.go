package web_question

import "github.com/backent/ai-golang/models"

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

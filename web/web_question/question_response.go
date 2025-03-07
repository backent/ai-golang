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
	Id              int64           `json:"id"`
	Name            string          `json:"name"`
	Language        string          `json:"language"`
	Chapter         string          `json:"chapter"`
	Amount          int             `json:"amount"`
	FileName        string          `json:"file_name"`
	StudentAttempts []StudentAttemp `json:"student_attempts"`
	Result          []ItemResult    `json:"result"`
}

type StudentAttemp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Score int16  `json:"score"`
}

func QuestionModelToQuestionGetByIdResponse(model models.Question) (QuestionGetByIdResponse, error) {
	var result Result
	err := json.Unmarshal([]byte(model.Result), &result)
	if err != nil {
		return QuestionGetByIdResponse{}, err
	}
	var studentAttempts []StudentAttemp
	if len(model.Exams) > 0 {
		for _, item := range model.Exams {
			studentAttempts = append(studentAttempts, examsToStudentAttemp(item))
		}
	} else {
		studentAttempts = make([]StudentAttemp, 0)
	}
	return QuestionGetByIdResponse{
		Id:              model.Id,
		Name:            model.Name,
		Language:        model.Language.String,
		Chapter:         model.Chapter.String,
		Amount:          model.Amount,
		FileName:        model.FileName,
		Result:          result.Result,
		StudentAttempts: studentAttempts,
	}, nil
}

func examsToStudentAttemp(model models.Exam) StudentAttemp {
	return StudentAttemp{
		Id:    model.Id,
		Name:  model.Username,
		Score: model.Score,
	}
}

type QuestionCheckMaterialResponse struct {
	Result bool `json:"result"`
}

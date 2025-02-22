package services_question

import (
	"context"

	"github.com/backent/ai-golang/web/web_question"
)

type QuestionServiceInterface interface {
	Create(ctx context.Context, request web_question.QuestionPostRequest) web_question.QuestionGetAllRequestItem
	GetAll(ctx context.Context) web_question.QuestionGetAllRequest
	GetById(ctx context.Context, id int) web_question.QuestionGetByIdResponse
	DeleteById(ctx context.Context, id int)
}

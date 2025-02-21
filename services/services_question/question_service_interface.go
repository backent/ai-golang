package services_question

import (
	"context"

	"github.com/backent/ai-golang/web/web_question"
)

type QuestionServiceInterface interface {
	Create(ctx context.Context, request web_question.QuestionPostRequest) web_question.Result
	GetAll(ctx context.Context) web_question.QuestionGetAllRequest
}

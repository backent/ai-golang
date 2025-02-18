package services_question

import "github.com/backent/ai-golang/web/web_question"

type QuestionServiceInterface interface {
	Create(request web_question.QuestionPostRequest) web_question.Result
}

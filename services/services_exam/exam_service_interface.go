package services_exam

import (
	"context"

	"github.com/backent/ai-golang/web/web_exam"
)

type ExamServiceInterface interface {
	GetExamByQuestionId(ctx context.Context, id int) web_exam.ExamGetByQuestionIdResponse
	GetExamById(ctx context.Context, id int) web_exam.ExamGetByQuestionIdResponse
	Submit(ctx context.Context, request web_exam.ExamSubmitRequest)
}

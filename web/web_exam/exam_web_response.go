package web_exam

import (
	"encoding/json"
	"time"

	"github.com/backent/ai-golang/models"
)

type ExamGetByQuestionIdResponse struct {
	Id        int64       `json:"id"`
	Questions []Questions `json:"questions"`
	Score     *int16      `json:"score"`
	ExamAt    *time.Time  `json:"exam_at"`
}

type Questions struct {
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	UserAnswer  *string  `json:"user_answer"`
	Answer      *string  `json:"answer"`
	Explanation *string  `json:"explanation"`
}

func ExamModelToExamGetByQuestionIdResponse(exam models.Exam) (ExamGetByQuestionIdResponse, error) {
	var response ExamGetByQuestionIdResponse
	var questions []Questions
	err := json.Unmarshal([]byte(exam.Submissions), &questions)
	if err != nil {
		return response, err
	}

	response.Id = exam.QuestionId
	response.Questions = questions
	response.Score = &exam.Score
	if !exam.ExamAt.Valid {
		for index := range response.Questions {

			response.Questions[index].UserAnswer = nil
			response.Questions[index].Answer = nil
			response.Questions[index].Explanation = nil
		}
		response.Score = nil
		response.ExamAt = nil
	} else {
		response.ExamAt = &exam.ExamAt.Time
	}
	return response, nil
}

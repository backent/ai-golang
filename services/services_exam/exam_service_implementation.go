package services_exam

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/rand"

	"github.com/backent/ai-golang/helpers"
	"github.com/backent/ai-golang/models"
	"github.com/backent/ai-golang/repositories/repositories_exam"
	"github.com/backent/ai-golang/repositories/repositories_question"
	"github.com/backent/ai-golang/web/web_exam"
	"github.com/backent/ai-golang/web/web_question"
)

type ExamServiceImplementation struct {
	*sql.DB
	repositories_exam.ExamRepositoryInterface
	repositories_question.RepositoryQuestionInterface
}

func NewExamServiceImplementation(db *sql.DB, exam_repo repositories_exam.ExamRepositoryInterface, question_repo repositories_question.RepositoryQuestionInterface) ExamServiceInterface {
	return &ExamServiceImplementation{DB: db, ExamRepositoryInterface: exam_repo, RepositoryQuestionInterface: question_repo}
}

func (implementation *ExamServiceImplementation) GetExamByQuestionId(ctx context.Context, id int) web_exam.ExamGetByQuestionIdResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	username, ok := ctx.Value(helpers.ContextKey("username")).(string)
	if !ok {
		panic("wrong username type")
	}

	exam, err := implementation.ExamRepositoryInterface.FindByQuestionIdANDUsername(ctx, tx, id, username)
	if err != nil {
		if err.Error() == "not found" {
			question, err := implementation.RepositoryQuestionInterface.GetById(ctx, tx, id)
			helpers.PanicIfError(err)

			var resultJson web_question.Result
			err = json.Unmarshal([]byte(question.Result), &resultJson)
			helpers.PanicIfError(err)

			resultJson.Result = shuffleOption(resultJson.Result)

			submissionByte, err := json.Marshal(resultJson.Result)
			helpers.PanicIfError(err)

			exam.Username = username
			exam.QuestionId = int64(id)
			exam.Submissions = string(submissionByte)

			exam, err = implementation.ExamRepositoryInterface.Create(ctx, tx, exam)
			helpers.PanicIfError(err)
		} else {
			panic(err)
		}
	}

	response, err := web_exam.ExamModelToExamGetByQuestionIdResponse(exam)
	helpers.PanicIfError(err)

	return response
}
func (implementation *ExamServiceImplementation) Submit(ctx context.Context, request web_exam.ExamSubmitRequest) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	username, ok := ctx.Value(helpers.ContextKey("username")).(string)
	if !ok {
		panic("wrong username type")
	}

	exam, err := implementation.ExamRepositoryInterface.FindByQuestionIdANDUsername(ctx, tx, int(request.Id), username)
	helpers.PanicIfError(err)
	var submissionJson []web_question.ItemResult

	err = json.Unmarshal([]byte(exam.Submissions), &submissionJson)
	helpers.PanicIfError(err)

	mappedSubmissionJson := make(map[string]web_question.ItemResult)
	for _, itemJson := range submissionJson {
		mappedSubmissionJson[itemJson.Question] = itemJson
	}

	var score int

	for index, submission := range request.Submissions {
		if targetQuestion, ok := mappedSubmissionJson[submission.Question]; ok {
			if submission.UserAnswer == targetQuestion.Answer {
				score++
			}
			request.Submissions[index].Answer = targetQuestion.Answer
			request.Submissions[index].Explanation = targetQuestion.Explanation
		}
	}

	submissionByte, err := json.Marshal(request.Submissions)
	helpers.PanicIfError(err)

	data := models.Exam{
		Username:    username,
		QuestionId:  request.Id,
		Submissions: string(submissionByte),
		Score:       int16(score),
	}

	implementation.ExamRepositoryInterface.UpdateByQuestionIdAndUsername(ctx, tx, data)
}

func shuffleOption(resultJson []web_question.ItemResult) []web_question.ItemResult {
	var answers = []string{"A", "B", "C", "D"}
	var mapAnswer = make(map[string]int)

	for index, answer := range answers {
		mapAnswer[answer] = index
	}

	for index, itemResult := range resultJson {
		var mapAnswerIndex int
		if numberIndex, ok := mapAnswer[itemResult.Answer]; ok {
			mapAnswerIndex = numberIndex
		}

		options := itemResult.Options
		rand.Shuffle(len(itemResult.Options), func(i, j int) {
			if i == mapAnswerIndex {
				mapAnswerIndex = j
			} else if j == mapAnswerIndex {
				mapAnswerIndex = i
			}
			itemResult.Options[i], itemResult.Options[j] = itemResult.Options[j], itemResult.Options[i]
		})
		itemResult.Options = options
		itemResult.Answer = answers[mapAnswerIndex]
		resultJson[index] = itemResult
	}

	return resultJson

}

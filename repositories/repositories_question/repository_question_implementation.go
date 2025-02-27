package repositories_question

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/backent/ai-golang/models"
)

type RepositoryQuestionImplementation struct {
}

func NewRepositoryQuestionImplementation() RepositoryQuestionInterface {
	return &RepositoryQuestionImplementation{}
}
func (implementation *RepositoryQuestionImplementation) Create(ctx context.Context, tx *sql.Tx, model models.Question) (models.Question, error) {
	query := fmt.Sprintf(`INSERT INTO %s (
	name,
	username,
	amount,
	gemini_file_uri,
	file_name,
	result
	) VALUES (?, ?, ?, ?, ?, ?)
	`, models.QuestionTable)

	result, err := tx.ExecContext(ctx, query, model.Name, model.Username, model.Amount, model.GeminiFileURI, model.FileName, model.Result)
	if err != nil {
		return model, err
	}

	model.Id, err = result.LastInsertId()
	if err != nil {
		return model, err
	}

	return model, nil
}

func (implementation *RepositoryQuestionImplementation) GetAll(ctx context.Context, tx *sql.Tx) ([]models.Question, error) {
	query := fmt.Sprintf("SELECT id, name, amount FROM %s", models.QuestionTable)
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []models.Question

	for rows.Next() {
		var item models.Question
		err = rows.Scan(&item.Id, &item.Name, &item.Amount)
		if err != nil {
			return nil, err
		}
		collections = append(collections, item)
	}

	return collections, nil
}

func (implementation *RepositoryQuestionImplementation) GetById(ctx context.Context, tx *sql.Tx, id int) (models.Question, error) {
	query := fmt.Sprintf("SELECT id, name, amount, file_name, result FROM %s WHERE id = ? LIMIT 1", models.QuestionTable)

	var model models.Question
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return model, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&model.Id, &model.Name, &model.Amount, &model.FileName, &model.Result)
		if err != nil {
			return model, err
		}
	} else {
		return model, errors.New("not found")
	}

	return model, nil
}

func (implementation *RepositoryQuestionImplementation) GetByIdWithExams(ctx context.Context, tx *sql.Tx, id int) (models.Question, error) {
	query := fmt.Sprintf("SELECT q.id, q.name, q.amount, q.file_name, q.result, e.id, e.username, e.score FROM %s q LEFT JOIN %s e ON q.id = e.question_id AND e.exam_at IS NOT NULL WHERE q.id = ?", models.QuestionTable, models.ExamTable)

	var model models.Question
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return model, err
	}
	defer rows.Close()

	mapQuestion := make(map[int64]*models.Question)

	for rows.Next() {
		var rowDataQuestion models.Question
		var rowDataExam models.NullAbleExam
		var questionData *models.Question
		err = rows.Scan(&rowDataQuestion.Id, &rowDataQuestion.Name, &rowDataQuestion.Amount, &rowDataQuestion.FileName, &rowDataQuestion.Result, &rowDataExam.Id, &rowDataExam.Username, &rowDataExam.Score)
		data, ok := mapQuestion[rowDataQuestion.Id]
		if !ok {
			mapQuestion[rowDataQuestion.Id] = &rowDataQuestion
			questionData = mapQuestion[rowDataQuestion.Id]
		} else {
			questionData = data
		}
		if rowDataExam.Id.Valid {
			questionData.Exams = append(questionData.Exams, models.Exam{
				Id:       rowDataExam.Id.Int64,
				Username: rowDataExam.Username.String,
				Score:    rowDataExam.Score.Int16,
			})
		}
		if err != nil {
			return models.Question{}, err
		}
	}

	return *mapQuestion[int64(id)], nil
}

func (implementation *RepositoryQuestionImplementation) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", models.QuestionTable)
	_, err := tx.ExecContext(ctx, query, id)

	return err
}

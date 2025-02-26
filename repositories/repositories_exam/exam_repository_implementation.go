package repositories_exam

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/backent/ai-golang/models"
)

type ExamRepositoryImplementation struct {
}

func NewExamRepositoryImplementation() ExamRepositoryInterface {
	return &ExamRepositoryImplementation{}
}

func (implementation *ExamRepositoryImplementation) Create(ctx context.Context, tx *sql.Tx, model models.Exam) (models.Exam, error) {
	query := fmt.Sprintf(`INSERT INTO %s (
	username,
	question_id,
	submissions,
	score
	) VALUES (?, ?, ?, ?)`, models.ExamTable)

	result, err := tx.ExecContext(ctx, query, model.Username, model.QuestionId, model.Submissions, model.Score)
	if err != nil {
		return model, err
	}
	model.Id, err = result.LastInsertId()
	if err != nil {
		return model, err
	}

	return model, nil
}
func (implementation *ExamRepositoryImplementation) FindByQuestionIdANDUsername(ctx context.Context, tx *sql.Tx, id int, username string) (models.Exam, error) {
	query := fmt.Sprintf("SELECT id, question_id, submissions, score, exam_at FROM %s WHERE question_id = ? AND username = ? LIMIT 1", models.ExamTable)

	var data models.Exam

	rows, err := tx.QueryContext(ctx, query, id, username)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&data.Id, &data.QuestionId, &data.Submissions, &data.Score, &data.ExamAt)
		if err != nil {
			return data, err
		}
	} else {
		return data, errors.New("not found")
	}

	return data, nil
}

func (implementation *ExamRepositoryImplementation) UpdateByQuestionIdAndUsername(ctx context.Context, tx *sql.Tx, model models.Exam) (models.Exam, error) {
	query := fmt.Sprintf("UPDATE %s SET score = ?, submissions = ?, exam_at = ? WHERE username = ? AND question_id = ?", models.ExamTable)

	_, err := tx.ExecContext(ctx, query, model.Score, model.Submissions, time.Now(), model.Username, model.QuestionId)
	if err != nil {
		return model, err
	}

	return model, nil

}

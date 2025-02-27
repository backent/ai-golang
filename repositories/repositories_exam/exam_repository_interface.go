package repositories_exam

import (
	"context"
	"database/sql"

	"github.com/backent/ai-golang/models"
)

type ExamRepositoryInterface interface {
	Create(ctx context.Context, tx *sql.Tx, model models.Exam) (models.Exam, error)
	FindByQuestionIdANDUsername(ctx context.Context, tx *sql.Tx, id int, username string) (models.Exam, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.Exam, error)
	UpdateByQuestionIdAndUsername(ctx context.Context, tx *sql.Tx, model models.Exam) (models.Exam, error)
}

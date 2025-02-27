package repositories_question

import (
	"context"
	"database/sql"

	"github.com/backent/ai-golang/models"
)

type RepositoryQuestionInterface interface {
	Create(ctx context.Context, tx *sql.Tx, model models.Question) (models.Question, error)
	GetAll(ctx context.Context, tx *sql.Tx) ([]models.Question, error)
	GetById(ctx context.Context, tx *sql.Tx, id int) (models.Question, error)
	GetByIdWithExams(ctx context.Context, tx *sql.Tx, id int) (models.Question, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int) error
}

package repositories_question

import (
	"context"
	"database/sql"

	"github.com/backent/ai-golang/models"
)

type RepositoryQuestionInterface interface {
	Create(ctx context.Context, tx *sql.Tx, model models.Question) (models.Question, error)
}

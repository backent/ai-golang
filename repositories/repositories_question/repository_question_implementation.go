package repositories_question

import (
	"context"
	"database/sql"
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

	for rows.Next() {
		err = rows.Scan(&model.Id, &model.Name, &model.Amount, &model.FileName, &model.Result)
		if err != nil {
			return model, err
		}
	}

	return model, nil
}

func (implementation *RepositoryQuestionImplementation) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", models.QuestionTable)
	_, err := tx.ExecContext(ctx, query, id)

	return err
}

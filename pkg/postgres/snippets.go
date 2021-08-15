package postgres

import (
	"database/sql"
	"knikoda/snippetbox/pkg/models"
	"time"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(expires string, title, content string) (int, error) {
	stmt := `INSERT INTO snippets (expires, title, content, created) VALUES($1, $2, $3, $4)`

	result, err := m.DB.Exec(stmt, expires, title, content, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get() (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

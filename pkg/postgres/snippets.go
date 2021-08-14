package postgres

import (
	"database/sql"
	"knikoda/snippetbox/pkg/models"
	"time"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(expires time.Time, title, content string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get() (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

package postgres

import (
	"database/sql"
	"errors"
	"knikoda/snippetbox/pkg/models"
	"time"
)

type SnippetModel struct {
	db *sql.DB
}

func New(db *sql.DB) *SnippetModel {
	return &SnippetModel{
		db: db,
	}
}

func (m *SnippetModel) Insert(expires string, title, content string) (int, error) {
	stmt := `INSERT INTO snippets (expires, title, content, created) VALUES($1, $2, $3, $4)`

	result, err := m.db.Exec(stmt, expires, title, content, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires 
		FROM snippets
		WHERE id = $1`

	row := m.db.QueryRow(stmt, id)

	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
		ORDER BY created DESC LIMIT 10`

	rows, err := m.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	return snippets, nil
}

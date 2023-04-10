package mysql

import (
	"database/sql"
	"errors"

	"github.com/petrostrak/clip-n-go/pkg/models"
)

// Defile a ClipModel type which wraps a sql.DB connection pool
type ClipModel struct {
	DB *sql.DB
}

func (m *ClipModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO clips (title, content, created, expires) 
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ClipModel) Get(id int) (*models.Clip, error) {
	stmt := `SELECT id, title, content, created, expires FROM clips
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	clip := &models.Clip{}
	err := m.DB.QueryRow(stmt, id).Scan(
		&clip.ID,
		&clip.Title,
		&clip.Content,
		&clip.Created,
		&clip.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return clip, nil
}

func (m *ClipModel) Latest() ([]*models.Clip, error) {
	return nil, nil
}

package mysql

import (
	"database/sql"

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
	return nil, nil
}

func (m *ClipModel) Latest() ([]*models.Clip, error) {
	return nil, nil
}

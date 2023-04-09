package mysql

import (
	"database/sql"

	"github.com/petrostrak/clip-n-go/pkg/models"
)

// Defile a ClipModel type which wraps a sql.DB connection pool
type ClipModel struct {
	DB *sql.DB
}

func (c *ClipModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (c *ClipModel) Get(id int) (*models.Clip, error) {
	return nil, nil
}

func (c *ClipModel) Latest() ([]*models.Clip, error) {
	return nil, nil
}

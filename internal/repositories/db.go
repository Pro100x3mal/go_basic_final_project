package repositories

import (
	"database/sql"
	"os"

	"github.com/Pro100x3mal/go_basic_final_project/internal/config"
	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '',
    title VARCHAR(256) NOT NULL DEFAULT '',
    comment TEXT NOT NULL DEFAULT '',
    repeat VARCHAR(128) NOT NULL DEFAULT ''
    );

CREATE INDEX scheduler_date ON scheduler(date);
`

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	var install bool
	if _, err := os.Stat(cfg.DBFile); os.IsNotExist(err) {
		install = true
	}

	db, err := sql.Open("sqlite", cfg.DBFile)
	if err != nil {
		return nil, err
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return nil, err
		}
	}
	return &Repository{db: db}, nil
}

func (r *Repository) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

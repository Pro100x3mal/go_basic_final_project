package repositories

import (
	"database/sql"
	"os"

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

func NewRepository() (*Repository, error) {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	var install bool
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		install = true
	}

	db, err := sql.Open("sqlite", dbFile)
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

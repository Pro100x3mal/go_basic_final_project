package repositories

import (
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) AddTask(task *models.Task) (int64, error) {
	var id int64

	if task == nil {
		return id, fmt.Errorf("database error: input is nil")
	}

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := r.db.Exec(query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)
	if err != nil {
		return id, fmt.Errorf("database error: failed to execute insert query: %w", err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("database error: failed to retrieve inserted task ID: %w", err)
	}

	return id, nil
}

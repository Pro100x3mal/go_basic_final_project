package repositories

import (
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) GetTask(id string) (*models.Task, error) {
	task := &models.Task{}

	query := `SELECT id, DATE, title, comment, repeat FROM scheduler WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks by ID %s: %w", id, err)
	}

	return task, nil
}

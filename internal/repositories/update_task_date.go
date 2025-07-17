package repositories

import (
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) UpdateTaskDate(task *models.Task) error {
	if task == nil {
		return fmt.Errorf("database error: cannot update task date: input is nil")
	}
	query := `UPDATE scheduler SET date=? WHERE id=?`
	res, err := r.db.Exec(query,
		task.Date,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("database error: failed to execute update query: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("database error: failed to retrieve affected rows: %w", err)
	}
	if count == 0 {
		return models.ErrTaskNotFound
	}

	return nil
}

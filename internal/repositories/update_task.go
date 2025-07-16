package repositories

import (
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) UpdateTask(task *models.Task) error {
	if task == nil {
		return fmt.Errorf("cannot update task: input is nil")
	}

	query := `UPDATE scheduler SET date=?, title=?, comment=?, repeat=? WHERE id=?`
	res, err := r.db.Exec(query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("no task found with ID: %s", task.ID)
	}

	return nil
}

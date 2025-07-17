package repositories

import (
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id=?`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("database error: failed to execute delete query: %w", err)
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

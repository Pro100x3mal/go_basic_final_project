package repositories

import "fmt"

func (r *Repository) DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id=?`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("no task deleted with ID: %s", id)
	}

	return nil
}

package repositories

import (
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) GetTasks(limit int) ([]*models.Task, error) {
	var out []*models.Task

	query := `SELECT * FROM scheduler ORDER BY date LIMIT ?`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := new(models.Task)
		err = rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)
		if err != nil {
			return nil, err
		}
		out = append(out, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return out, err
}

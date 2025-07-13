package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) GetTasks(limit int) ([]*models.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY DATE LIMIT ?`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	defer rows.Close()

	return scanTasks(rows)
}

func (r *Repository) GetTasksByDate(date string, limit int) ([]*models.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE DATE = ? LIMIT ?`
	rows, err := r.db.Query(query, date, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks by date %s: %w", date, err)
	}
	defer rows.Close()

	return scanTasks(rows)
}

func (r *Repository) GetTasksByKeyword(search string, limit int) ([]*models.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY DATE LIMIT ?`
	rows, err := r.db.Query(query, "%"+search+"%", "%"+search+"%", limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks by keyword %q: %w", search, err)
	}
	defer rows.Close()

	return scanTasks(rows)
}

func scanTasks(rows *sql.Rows) ([]*models.Task, error) {
	var out []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		out = append(out, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return out, nil
}

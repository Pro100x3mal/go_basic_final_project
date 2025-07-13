package repositories

import (
	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (r *Repository) AddTask(task *models.Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := r.db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

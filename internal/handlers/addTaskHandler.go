package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Pro100x3mal/go_basic_final_project/internal/models"
)

func (th *TaskHandler) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &task)
	if err != nil {
		writeJson(w, err)
		return
	}

	if task.Title == "" {
		writeJson(w, errors.New("title is required"))
		return
	}

	id, err := th.taskService.CreateTask(&task)
	if err != nil {
		writeJson(w, err)
		return
	}

	writeJson(w, strconv.Itoa(int(id)))
}

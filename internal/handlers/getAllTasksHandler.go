package handlers

import "net/http"

func (th *TaskHandler) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	limit := 50
	tasks, err := th.taskService.GetAllTasks(limit)
	if err != nil {
		writeJson(w, err)
	}
	writeJson(w, tasks)
}

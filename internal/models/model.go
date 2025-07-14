package models

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type RespID struct {
	ID string `json:"id"`
}

type RespError struct {
	Error string `json:"error"`
}

type RespOk struct{}

type RespTasks struct {
	Tasks []*Task `json:"tasks"`
}

type Password struct {
	Password string `json:"password"`
}

type RespToken struct {
	Token string `json:"token"`
}

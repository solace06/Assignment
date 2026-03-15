package internal

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
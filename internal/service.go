package internal

func NewTask(task CreateTaskRequest) (*Task, error) {
	var newTask Task
	var status string

	id := generateUUID()

	if task.Status == "" {
		status = StatusTodo
	} else {
		status = task.Status
	}

	newTask = Task{
		ID:     id,
		Title:  task.Title,
		Status: status,
	}

	insertedTask, err := InsertTask(newTask)
	if err != nil {
		return nil, err
	}

	return insertedTask, nil
}

func GetAllTasks() ([]Task, error) {
	return FetchTasks()
}

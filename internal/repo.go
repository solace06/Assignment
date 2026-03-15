package internal

var tasks = []Task{
	{ID: generateUUID(), Title: "Learn Go", Status: "todo"},
	{ID: generateUUID(), Title: "Write APIs", Status: "in_progress"},
}

func InsertTask(task Task) (*Task, error) {
	tasks = append(tasks, task)

	return &task, nil
}

func FetchTasks() ([]Task, error) {
	return tasks, nil
}

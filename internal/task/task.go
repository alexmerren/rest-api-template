package task

// Task struct is used to encapsulate functions
type Task struct {
	ID      int
	Content string
}

// NewTask function is called to return a new task with specified information
func NewTask(id int, content string) *Task {
	return &Task{
		ID:      id,
		Content: content,
	}
}

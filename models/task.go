package tasks

// Task
type Task struct {
	ID          string
	Name        string
	Description string
	IsDone      bool
}

// NewTask
func NewTask(id, name, description, isDone string) *Task {
	return &Task{
		ID:          id,
		Name:        name,
		Description: description,
		IsDone:      false,
	}
}

// TaskList
type TaskList struct {
	Items []Task
}

// TaskInfo
func (t Task) TaskInfo() string {
	return "Task: " + t.Name + " - " + t.Description + " "
}

// AddTask
func (t Task) AddTask(l TaskList) {
	l.Items = append(l.Items, t)
}

// CompleteTask
func (t *Task) Do() {
	t.IsDone = true
}

// ListTasks
func (t TaskList) ListTasks() {
	for _, task := range t.Items {
		if !task.IsDone {
			task.TaskInfo()
		}
	}
}

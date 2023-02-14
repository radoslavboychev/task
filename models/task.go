package models

import "fmt"

// Task
type Task struct {
	ID     string
	Name   string
	IsDone bool
}

// NewTask
func NewTask(id, name string, isDone bool) *Task {
	return &Task{
		ID:     id,
		Name:   name,
		IsDone: false,
	}
}

// TaskList
type TaskList struct {
	Items []Task
}

// TaskInfo
func (t Task) TaskInfo() string {
	fmt.Println("Task: " + t.ID + " " + t.Name)
	return "Task: " + t.ID + " " + t.Name
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

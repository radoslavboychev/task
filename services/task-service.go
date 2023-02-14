package services

import (
	"errors"
	"fmt"

	"github.com/radoslavboychev/task/models"
	"github.com/radoslavboychev/task/repo"
)

type TaskService struct {
	repository repo.TaskRepo
	taskList   models.TaskList
}

func NewTaskService(repository repo.TaskRepo, taskList models.TaskList) *TaskService {
	return &TaskService{
		repository: repository,
		taskList:   taskList,
	}
}

func (t TaskService) ListTasks() error {
	for _, task := range t.taskList.Items {
		fmt.Printf("task: %v\n", task)
	}

	err := t.repository.ViewDB("tasks")
	if err != nil {
		return err
	}
	return nil
}

// DoTask
func (t *TaskService) DoTask(id string) error {

	for _, task := range t.taskList.Items {
		if task.ID == id {
			task.IsDone = true
			err := t.repository.DoTask(id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddTask
func (t *TaskService) AddTask(in models.Task) error {
	for _, task := range t.taskList.Items {
		if task.ID == in.ID {
			return errors.New("ID is taken")
		}
	}

	in.AddTask(t.taskList)
	err := t.repository.CreateTask(in)
	if err != nil {
		return err
	}

	return nil
}

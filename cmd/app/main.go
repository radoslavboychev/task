package main

import (
	"log"

	"github.com/radoslavboychev/task/models"
	"github.com/radoslavboychev/task/repo"
)

func main() {
	// set up test data
	task1 := models.NewTask("1", "do dishes", false)
	task2 := models.NewTask("2", "do laundry", false)
	task3 := models.NewTask("3", "clean closet", false)
	task4 := models.NewTask("4", "wash vests", false)
	task5 := models.NewTask("5", "find keys", false)
	task6 := models.NewTask("6", "play football", false)

	// add all tasks to a list
	list := models.TaskList{}
	list.Items = append(list.Items, *task1, *task2, *task3, *task4, *task5, *task6)

	// initiate DB connection
	db, err := repo.InitDB("tasks.db", 0600)
	if err != nil {
		return
	}
	log.Println("DB initialized")

	// create new repository
	r := repo.NewTaskRepo(db)

	// create a new key-value bucket
	r.CreateBucket("tasks")

	// add tasks to the service
	// taskManager := services.NewTaskService(r, list)
	// for _, i := range list.Items {
	// 	taskManager.AddTask(i)
	// }

	err = r.ViewDB("tasks")
	if err != nil {
		log.Println(err)
		return
	}
	// err = taskManager.DoTask("1")
	// if err != nil {
	// 	return
	// }

	// taskManager.ListTasks()

	// defer closing the DB connection
	defer r.DB.Close()

}

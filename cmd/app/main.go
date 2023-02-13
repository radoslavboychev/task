package main

import (
	"flag"
	"log"

	"github.com/radoslavboychev/task/models"
	"github.com/radoslavboychev/task/repo"
)

func main() {
	// set up test data
	task1 := models.NewTask("15943", "Do dishes", false)
	task2 := models.NewTask("15944", "Do laundry", false)
	task3 := models.NewTask("15945", "Clean closet", false)
	task4 := models.NewTask("15946", "Wash vests", false)
	task5 := models.NewTask("15947", "Find keys", false)
	task6 := models.NewTask("15948", "Play football", false)

	list := models.TaskList{}
	list.Items = append(list.Items, *task1, *task2, *task3, *task4, *task5, *task6)

	db, err := repo.InitDB("tasks.db", 0600)
	log.Println("DB initiated")
	if err != nil {
		return
	}

	log.Println("Database loaded")

	// create new repository
	r := repo.NewTaskRepo(db)

	// create a new key-value bucket
	r.CreateBucket("tasks")

	// defer closing the DB connection
	defer r.DB.Close()

	// create and add tasks to BoltDB
	for _, t := range list.Items {
		err = r.CreateTask(t)
		if err != nil {
			return
		}
	}

	// print all tasks
	err = r.ViewDB("tasks")
	if err != nil {
		return
	}

}

//
type generalFlags struct {
	Name string
}

//
func (c *generalFlags) Flags() *flag.FlagSet {

	fs := flag.NewFlagSet("helpFS", flag.ContinueOnError)
	fs.String("help", "help", "use to show all commands")
	return fs
}

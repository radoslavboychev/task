package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/radoslavboychev/task/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

// addCmd
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a command to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, "")
		_, err := db.CreateTask(task)
		if err != nil {
			log.Printf("Error occurred: %v", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list \n", task)
	},
}

// doCmd
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Printf("Error occurred: %v", err)
				return
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				log.Printf("Failed to complete task %v: %v", id, err)
				return
			}
		}
	},
}

// listCmd
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			log.Printf("Error occurred: %v", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

// init
func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(doCmd)
	RootCmd.AddCommand(listCmd)
}

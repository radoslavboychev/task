package main

import (
	"fmt"
	"os"

	cmd "github.com/radoslavboychev/task/cobra"
	"github.com/radoslavboychev/task/db"
)

func main() {
	// home, _ := homedir.Dir()
	dbPath := "./tasks.db"
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

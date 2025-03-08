package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/isoment/taskmaster/cmd"
	"github.com/isoment/taskmaster/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))

	fmt.Println("Database connection established")

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

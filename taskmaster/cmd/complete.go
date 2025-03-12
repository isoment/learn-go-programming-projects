package cmd

import (
	"fmt"

	"github.com/isoment/taskmaster/db"
	"github.com/spf13/cobra"
)

var completeCommand = &cobra.Command{
	Use:   "complete",
	Short: "Show all of the completed tasks in the last 24hrs",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.CompletedTasks()
		if err != nil {
			fmt.Println("There was an error getting tasks: ", err)
		}

		if len(tasks) == 0 {
			fmt.Println("You did not complete any tasks in the past 24 hours.")
		}

		fmt.Println("You completed the following tasks in the last 24hrs...")
		for _, t := range tasks {
			formattedTime := t.Value.CompletedAt.Format("2006-01-02 15:04:05")
			fmt.Printf("You completed \"%s\" on: \"%s\"\n", t.Value.Description, formattedTime)
		}
	},
}

func init() {
	RootCmd.AddCommand(completeCommand)
}

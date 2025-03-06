package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

/*
The init() function gets executed before the main() application function. There can be multiple init()
functions in a package. If that is the case they get executed in the order they occur. In this case we
are adding the add command to the root command
*/
func init() {
	RootCmd.AddCommand(addCmd)
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "TaskTracker",
	Short: "CLI application to track you'r daily tasks",
	Long: `CLI application to track you'r daily tasks
	Every task contains: Creator of task (--user/-u default "Me"), task (--task/-t default "None") and progress-mark (--mark/-m default "in progress").
	commands:
		add		command to create new task
			Example: add --user Me --task "Go outside" 

		delete	command to delete you'r --tasks or users
			Example: delete --user Me --task "Go outside"

		mark	command to change progress-mark
			Example: mark --user Me --task "Go outside" --mark Done
			
		show	command to show tasks
			Example show --user Me
	
	`,

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



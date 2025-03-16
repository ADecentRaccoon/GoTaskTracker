package cmd

import (
	"TaskTracker/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showMark string
var showUsername string
var showTaskname string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Command to show you'r tasks",
	Long: `This command show you'r tasks from json file
		Usage: show --user {Username default "all"} --task {Name of task default "all"} --mark {Mark of task default "all"}
		
		Examples: 
			show			- shows all tasks, users, and progress marks
			show -u Me		- shows all tasks of user "Me"
			show -m Done	- shows all tasks with progress mark "Done"
			`,
	Run: func(cmd *cobra.Command, args []string) {
		show(&showMark, &showUsername, &showTaskname, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&showUsername, "user", "u", "all", "docs")
	showCmd.Flags().StringVarP(&showMark, "mark", "m", "all", "docs")
	showCmd.Flags().StringVarP(&showTaskname, "tasks", "t", "all", "docs")
}

func show(mark *string, user *string, taskToShow *string, filename string) {
	tasks := pkg.LoadTask(filename)

	nameCounter := 0

	regName := func(name *string, nameCounter *int) {
		if *nameCounter > 0 {
			*nameCounter -= 1
			fmt.Println("\t", *name)
		}
	}

	if *user != "all" {
		if *taskToShow != "all" {
			if *mark != "all" {
				if tempAns := tasks[*user][*taskToShow]; tempAns == *mark {
					nameCounter += 1
					regName(user, &nameCounter)
					fmt.Println(tempAns)
				} else {
					os.Exit(0)
				}
			} else {
				for iterTask, iterMark := range tasks[*user] {
					if (iterTask == *taskToShow || iterTask == "all") && (iterMark == *mark || *mark == "all") {
						fmt.Println(tasks[*user][iterTask])
					}
				}
			}
		}
	} else {
		for iterUser, iterTasks := range tasks {
			nameCounter += 1
			for iterTask, iterMarks := range iterTasks {
				if (iterTask == *taskToShow || *taskToShow == "all") && (*mark == "all" || *mark == iterMarks) {
					regName(&iterUser, &nameCounter)
					fmt.Printf("|%-25s|%-25s|\n", iterTask, iterMarks)
				}
			}
		}
	}
}

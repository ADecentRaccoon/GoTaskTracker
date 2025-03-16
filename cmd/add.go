package cmd

import (
	"TaskTracker/pkg"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)
var addMark string
var addUsername string
var addTaskname string


var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add tasks",
	Long:  
	`This command add tasks to you'r json file
		Usage: add --user {Username default "Me"} --task {Name of task default "None"} --mark {Mark of task default "in progress"}
		
		Examples: 
			add --user Me --task "go outside" --mark "in progress"
			add -u Me -t "go outside"
			add -t "go outside"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(&addUsername, &addTaskname, &addMark, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addUsername, "username", "u", "Me", "Flag to add username to new task")
	addCmd.Flags().StringVarP(&addTaskname, "taskname", "t", "None", "Flag to add tasks to new task")
	addCmd.Flags().StringVarP(&addMark, "mark", "m", "in progress", "Flag to add progress-mark to new task")
}

func addTask(user *string, task *string, mark *string, filename string) {
	tasks := pkg.LoadTask(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if userTasks, exist := tasks[*user]; exist {
		userTasks[*task] = *mark
	} else {
		tasks[*user] = make(map[string]string)
		tasks[*user][*task] = *mark
	}
	jsoned, parceErr := json.Marshal(tasks)
	if parceErr != nil {
		panic(parceErr)
	}
	file.Write(jsoned)
}

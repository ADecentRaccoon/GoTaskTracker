/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"TaskTracker/pkg"
	"encoding/json"
	"os"
)

var username string
var taskname string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add tasks",
	Long: `Thise command add tasks to you'r json file`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(&username, &taskname, &mark, "data.json")
	},
}



func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&username, "username", "u", "Me", "Flag to add username to new task")
	addCmd.Flags().StringVarP(&taskname, "taskname", "t", "None", "Flag to add taks to new task")
	addCmd.Flags().StringVarP(&mark, "mark", "m", "in progress", "Flag to add progress-mark to new task")
}


func addTask(user *string, task *string, mark *string, filename string) {
	tasks := pkg.LoadTask(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if userTasks, exist := tasks[*user]; exist == true{
		userTasks[*task] = *mark
	} else{
		tasks[*user] = make(map[string]string)
		tasks[*user][*task] = *mark
	}
	tasks[*user] = map[string]string{*task: *mark}
	jsoned, parceErr := json.Marshal(tasks)
	if parceErr != nil {
		panic(parceErr)
	}
	file.Write(jsoned)
}
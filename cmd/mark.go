/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"TaskTracker/pkg"
	"os"
)

var mark string


// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		markTaks(&username, &taskname, &mark, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
	markCmd.Flags().StringVarP(&username, "user", "u", "all", "")
	markCmd.Flags().StringVarP(&taskname, "task", "t", "all", "")
	markCmd.Flags().StringVarP(&mark, "mark", "m", "Done", "")
}

func markTaks(user *string, taskToDMark *string, mark *string, filename string) {
	if *user == "all" && *taskToDMark == "all" { // in case of all progress-markes should be change 
		fmt.Printf("\tAre you sure that you want to mark all tasks as %s? [y/N]\n", *mark)
		ans := ""
		fmt.Scanln(&ans)
		if ans == "Y" || ans == "yes" || ans == "Yes"|| ans == "y" {
			tasks := pkg.LoadTask(filename)
			for iterUser, usersTasks := range tasks{
				for userTask, _ := range usersTasks{
					tasks[iterUser][userTask] = *mark
				}
			}
		}
	}

	tasks := pkg.LoadTask("data.json")

	if *user != "all" && *taskToDMark != "all"{ // fast progress-mark changing O(1)
		tasks[*user][*taskToDMark] = *mark
		} else if *user != "all"{ // iteration thrue tasks. Much slower (O(n) n - quantity of tasks)
			for iterTask, _ := range tasks[*user]{
				tasks[*user][iterTask] = *taskToDMark
			} 
		} else if *taskToDMark != "all"{ // iteration thrue users. O(n) n - quantity of users
			for iterUser, _ := range tasks{
				tasks[iterUser][*taskToDMark] = *taskToDMark
			}
		}
	
	
	
		
	jsoneded, errParce := json.Marshal(tasks) // marshaling changes
	if errParce != nil{
		panic(errParce)
	}

	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0666) // saving changes
	if err != nil{
		panic(err)
	}
	defer file.Close()
	_, errWrite := file.Write(jsoneded)
		if errWrite != nil{
			panic(errWrite)
		}
}
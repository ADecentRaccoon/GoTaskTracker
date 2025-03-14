/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TaskTracker/pkg"
	"fmt"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		show(&username, &taskname, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&username, "user", "u", "all", "docs")
	showCmd.Flags().StringVarP(&taskname, "tasks", "t", "all", "docs")
}


func show(user *string, taskToShow *string, filename string){
	tasks := pkg.LoadTask(filename)

	if *user != "all"{
		fmt.Println("\t", *user)
		for iterTasks, iterMarks := range tasks[*user]{
			if iterTasks == *taskToShow || iterTasks == "all"{
				fmt.Println("\t\t", iterTasks, "\t", iterMarks)
			}
		}
	} else{
		for iterUser, iterTasks := range tasks{
			fmt.Println("\t", iterUser)
			for iterTask, iterMarks := range iterTasks{
				if iterTask == *taskToShow || *taskToShow == "all"{
					fmt.Println("\t\t", iterTask, "\t", iterMarks)
				}
			}
		}
	}



}
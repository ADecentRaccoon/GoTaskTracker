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
	if *user == "all" && *taskToDMark == "all" {
		fmt.Printf("\tAre you sure that you want to mark all tasks as %s? [y/N]\n", *mark)
		ans := ""
		fmt.Scanln(&ans)
		if ans == "Y" || ans == "yes" || ans == "Yes"|| ans == "y" {
			tasks := pkg.LoadTask(filename)
			for _, task := range tasks{
				task[2] = *mark
			}
		}
	}
	tasks := pkg.LoadTask("data.json")
	file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	for i := range tasks{
		if (*user == username || *user == "all") && (*taskToDMark == taskname || *taskToDMark == "all"){
			tasks[i][2] = *mark
		}
	}
	jsoneded, errParce := json.Marshal(tasks)
	if errParce != nil{
		panic(errParce)
	}
	_, errWrite := file.Write(jsoneded)
		if errWrite != nil{
			panic(errWrite)
		}
}
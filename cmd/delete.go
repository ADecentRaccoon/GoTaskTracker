/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"TaskTracker/pkg"
	"github.com/spf13/cobra"
	"os"
	"encoding/json"
	"fmt"
)

var user string
var taskToDelete string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "commant to delete you'r tasks",
	Long: ``,// finish later
	Run: func(cmd *cobra.Command, args []string) {
		deleteTaks(&user, &taskToDelete, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&user, "user", "u", "all", "Select user to delete task")
	deleteCmd.Flags().StringVarP(&taskToDelete, "task", "t", "all", "Select task to delete")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func deleteTaks(user *string, taskToDelete *string, filename string) {
	if *user == "all" && *taskToDelete == "all" {
		fmt.Print("\tAre you sure that you want to delete all tasks? [y/N]\t")
		ans := ""
		fmt.Scanln(&ans)
		if ans == "Y" || ans == "yes" || ans == "Yes"|| ans == "y" {
			os.Remove(filename)
			os.Create(filename)
		}
	}
		tasks := pkg.LoadTask(filename)
		
		if *user != "all" && *taskToDelete != "all"{
			delete(tasks[*user], *taskToDelete)
		} else if *user != "all"{
			for iterTask, _ := range tasks[*user]{
				delete(tasks[*user], iterTask)
			}
		}

		jsoneded, errParce := json.Marshal(tasks)
		if errParce != nil{
			panic(errParce)
		}
		
		file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil{
			panic(err)
		}

		_, errWrite := file.Write(jsoneded)
		if errWrite != nil{
			panic(errWrite)
		}
		defer file.Close()
	}

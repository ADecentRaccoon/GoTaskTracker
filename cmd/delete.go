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

var deleteUser string
var deleteTask string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "command to delete you'r tasks",
	Long: `This command deletes tasks or users
		Usage: delete --user {Username default "all"} --task {Name of task default "all"}
		Examples: 
			delete --user Me --task "go outside"					- deletes task "go outside" from user "Me"
			delete -u Me											- deletes all tasks of user "Me"
			delete -m "Blocked"										- deletes all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteTaks(&deleteUser, &deleteTask, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteUser, "user", "u", "all", "Select user to delete task")
	deleteCmd.Flags().StringVarP(&deleteTask, "task", "t", "all", "Select task to delete")
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

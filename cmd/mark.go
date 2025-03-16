package cmd

import (
	"TaskTracker/pkg"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var markUsername string
var markTaskname string
var markMark string


var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "A brief description of your command",
	Long: `This command changes task's progress-marks
		Usage: mark --user {Username default "all"} --task {Name of task default "all"} --mark {Mark of task default "in progress"}
		
		Examples: 
			mark --user Me --task "go outside" --mark "Blocked"	- marks task "go outside" of user "Me" as "Blocked"
			mark -u Me -t "go outside"							- marks task "go outside" of user "Me" as "Done"
			mark -u Me --mark "Blocked"							- marks all tasks of user "Me" as "Blocked"
			mark -m "Blocked"									- marks all tasks as "Blocked"
			`,
	Run: func(cmd *cobra.Command, args []string) {
		markTaks(&markUsername, &markTaskname, &markMark, "data.json")
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
	markCmd.Flags().StringVarP(&markUsername, "user", "u", "all", "")
	markCmd.Flags().StringVarP(&markTaskname, "task", "t", "all", "")
	markCmd.Flags().StringVarP(&markMark, "mark", "m", "Done", "")
}

func markTaks(user *string, taskToDMark *string, lmark *string, filename string) {

	saveChangings := func(tasks map[string]map[string]string) {
		jsoneded, errParce := json.Marshal(tasks) // marshaling changes
		if errParce != nil {
			panic(errParce)
		}

		file, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0666) // saving changes
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, errWrite := file.Write(jsoneded)
		if errWrite != nil {
			panic(errWrite)
		}
		os.Exit(0)
	}

	if *user == "all" && *taskToDMark == "all" { // in case of all progress-markes should be change
		fmt.Printf("\tAre you sure that you want to mark all tasks as %s? [y/N]\n", *lmark)
		ans := ""
		fmt.Scanln(&ans)
		if ans == "Y" || ans == "yes" || ans == "Yes" || ans == "y" {
			tasks := pkg.LoadTask(filename)
			for iterUser, usersTasks := range tasks {
				for userTask, _ := range usersTasks {
					tasks[iterUser][userTask] = *lmark
				}
			}
			saveChangings(tasks)
		}
	}

	tasks := pkg.LoadTask("data.json")

	if *user != "all" && *taskToDMark != "all" { // fast progress-mark changing O(1)
		tasks[*user][*taskToDMark] = *lmark
		saveChangings(tasks)
	} else if *user != "all" { // iteration through tasks. Much slower (O(n) n - quantity of tasks)
		for iterTask, _ := range tasks[*user] {
			tasks[*user][iterTask] = *lmark
		}
		saveChangings(tasks)
	} else if *taskToDMark != "all" { // iteration through users. O(n) n - quantity of users
		for iterUser, _ := range tasks {
			tasks[iterUser][*taskToDMark] = *taskToDMark
		}
		saveChangings(tasks)
	}

}

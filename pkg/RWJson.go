package pkg

import (
	"encoding/json"
	"os"
)


func LoadTask(filename string) [][3]string {
	tasks, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		os.Create(filename)
		return [][3]string{}
	}
	if len(tasks) == 0 {
		return [][3]string{}
	}
	var answer [][3]string
	parceErr := json.Unmarshal(tasks, &answer)
	if parceErr != nil {
		panic(parceErr)
	}
	return answer
}

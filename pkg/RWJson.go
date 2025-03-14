package pkg

import (
	"encoding/json"
	"os"
)


func LoadTask(filename string) map[string]map[string]string {
	tasks, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		os.Create(filename)
		return map[string]map[string]string{}
	}
	if len(tasks) == 0 {
		return map[string]map[string]string{}
	}
	var answer map[string]map[string]string
	parceErr := json.Unmarshal(tasks, &answer)
	if parceErr != nil {
		panic(parceErr)
	}
	return answer
}

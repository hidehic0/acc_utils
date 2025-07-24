package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	types "hidehic0/acc_utils/internal/type"
)

func GetInfomation() types.Infomation {
	bytes, err := os.ReadFile("./contest.acc.json")

	if err != nil {
		fmt.Println("file read error")
		log.Fatal(err)
		os.Exit(256)
	}

	var res types.Infomation

	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Println("json parse error")
		log.Fatal(err)
		os.Exit(256)
	}

	return res
}

func GetTasks() []string {
	res := []string{}

	for _, task := range GetInfomation().Tasks {
		res = append(res, task.Directory.Path)
	}

	return res
}

func GetTaskInfomation() map[string]types.TaskInfomation {
	res := make(map[string]types.TaskInfomation)

	infomation := GetInfomation()

	for _, task := range infomation.Tasks {
		res[task.Directory.Path] = task
	}

	return res
}

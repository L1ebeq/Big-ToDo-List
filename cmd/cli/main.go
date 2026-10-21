package main

import (
	"bufio"
	"fmt"
	"os"
	"task-manager/internal/models"
	"task-manager/internal/variables"
	"task-manager/storage"
)

func main() {
	fileName := "data.json"
	tasks, err := storage.ReadData(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []models.Task{}
		}
	} else {
		fmt.Println("Read file error")
	}
	variables.TasksSlice = tasks
	storage.SaveData(fileName, variables.TasksSlice)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		models.ShowMenu()
		var s string
		scanner.Scan()
		s = scanner.Text()
		end := models.MenuProcessing(s, &variables.TasksSlice, scanner)
		if end {
			return
		}
	}

}

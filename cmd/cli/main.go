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
	//models.AddOneTaks(variables.Task1, &variables.TasksSlice)
	//models.AddOneTaks(variables.Task2, &variables.TasksSlice)
	fileName := "jsonFile"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("File creation error")
	}
	storage.SaveData(file, variables.TasksSlice)
	for {
		models.ShowMenu()
		var s string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
		//fmt.Scan(&s)
		end := models.MenuProcessing(s, &variables.TasksSlice, scanner)
		if end {
			return
		}
	}

}

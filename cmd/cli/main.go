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
	fileName := "jsonFile"
	_, err := os.Create(fileName)
	if err != nil {
		fmt.Println("File creation error")
		return
	}
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

package main

import (
	"bufio"
	"os"
	"task-manager/internal/models"
	"task-manager/internal/variables"
)

func main() {
	//models.AddOneTaks(variables.Task1, &variables.TasksSlice)
	//models.AddOneTaks(variables.Task2, &variables.TasksSlice)
	for {
		models.ShowMenu()
		var s string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
		//fmt.Scan(&s)
		end := models.MenuProcessing(s, &variables.TasksSlice)
		if end {
			return
		}
	}
}

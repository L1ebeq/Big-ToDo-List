package main

import (
	"fmt"
	"task-manager/internal/models"
	"task-manager/internal/variables"
)

func main() {
	models.AddOneTaks(variables.Task1, &variables.TasksSlice)
	models.AddOneTaks(variables.Task2, &variables.TasksSlice)
	for {
		fmt.Println("Выберите команду")
		fmt.Println("1. Показать задачу")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Изменить статус")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")
		var s, name string
		fmt.Println("Введите команду")
		fmt.Scan(&s)
		switch {
		case s == "1":
			models.ListTasks(variables.TasksSlice)
		case s == "2":
			fmt.Println("Введите название задачи")
			fmt.Scan(&name)
			models.Add(name, &variables.TasksSlice)
		case s == "3":
			var id int
			fmt.Println("Введите ID")
			fmt.Scan(&id)
			models.Done(id, variables.TasksSlice)
		case s == "4":
			var id int
			fmt.Println("Введите ID")
			fmt.Scan(&id)
			models.Delete(id, &variables.TasksSlice)
		case s == "5":
			return
		}
	}
}

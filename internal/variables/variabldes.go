package variables

import (
	"task-manager/internal/models"
)

var Task1 = models.Task{
	ID:        9,
	Title:     "выучить структуры",
	Completed: false,
}

var Task2 = models.Task{
	ID:        2,
	Title:     "выучить слайсы",
	Completed: false,
}
var Task3 = models.Task{
	ID:        3,
	Title:     "выучить функции",
	Completed: false,
}
var TasksSlice []models.Task

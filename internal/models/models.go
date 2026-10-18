package models

import (
	"bufio"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	DeadLine    time.Time `json:"dead_line"`
	Tags        []string  `json:"tags"`
}

func AddOneTaks(a Task, b *[]Task) {

	*b = append(*b, a)
}

func ListTasks(s []Task) { //Вывод задач

	for _, v := range s {
		fmt.Println("ID:", v.ID)
		fmt.Println(v.Title)
		fmt.Println(v.Description)
		if v.Completed == true {
			fmt.Println("Статус: Выполнено")
		} else {
			fmt.Println("Статус: Не выполнено")
		}
		fmt.Println(v.CreatedAt.Format("2006-01-02 15:04:05"))
		GetDeadlineStatus(v.DeadLine)
	}

}

func Add(str, des string, prio int, tm time.Time, s *[]Task) { //Создание задачи
	var l int
	l = len(*s)
	var p string
	switch {
	case prio == 1:
		p = "High"
	case prio == 2:
		p = "Medium"
	case prio == 3:
		p = "Low"
	default:
		fmt.Println("Incorrect priority")
		return
	}
	var t = Task{
		ID:          l + 1,
		Title:       str,
		Completed:   false,
		CreatedAt:   time.Now(),
		Description: des,
		Priority:    p,
		DeadLine:    tm,
	}
	*s = append(*s, t)
	fmt.Println(t)
}

func Done(id int, s []Task) (error, bool) { //Смена статуса на Выполнено
	found := false
	for i, v := range s {
		if v.ID == id {
			(s)[i].Completed = true
			found = true
		}
	}
	if found == true {
		return nil, found
	} else {
		err := errors.New("Invalid ID")
		fmt.Println("Error:", err)
		return err, found
	}
}

func Delete(id int, s *[]Task) (error, bool) { //Удаление задачи
	found := false
	for i, v := range *s {
		if v.ID == id {
			(*s)[i] = (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			found = true
		}
	}
	if found == true {
		return nil, found
	} else {
		err := errors.New("Invalid ID")
		fmt.Println("Error:", err)
		return err, found
	}
}

func ShowMenu() { //Отображение меню
	fmt.Println("Выберите команду")
	fmt.Println("1. Показать задачу")
	fmt.Println("2. Добавить задачу")
	fmt.Println("3. Изменить статус")
	fmt.Println("4. Удалить задачу")
	fmt.Println("5. Найти задачу")
	fmt.Println("6. Отсортировать")
	fmt.Println("7. Выйти")
	fmt.Println("Введите команду")
}

func TaskSearch(str string, s []Task) { // Поиск задачи по имени
	found := false
	for _, v := range s {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(str)) {
			fmt.Println(v.Title, v.Description)
			found = true
		} else {
			if strings.Contains(strings.ToLower(v.Description), strings.ToLower(str)) {
				fmt.Println(v.Title, v.Description)
				found = true
			}
		}
	}
	if found != true {
		fmt.Println("String doesn't found")
	}
}

func MenuProcessing(s string, t *[]Task, scanner *bufio.Scanner) bool { //Обработка команд
	end := false
	var name, descript string
	var timeInput string
	switch {
	case s == "1":
		ListTasks(*t)
	case s == "2":
		fmt.Println("Введите название задачи")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("Введите описание задачи")
		scanner.Scan()
		descript = scanner.Text()
		fmt.Println("Выберите приоритет")
		fmt.Println("1.High 2. Medium 3. Low")
		scanner.Scan()
		prio, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Нужно ввести число")
			return false
		}
		fmt.Println("Укажите дэдлайн")
		scanner.Scan()
		timeInput = scanner.Text()
		tm, err := time.Parse("2006-01-02 15:04:05", timeInput)
		if err != nil {
			fmt.Println("Incorrect time format")
		}
		Add(name, descript, prio, tm, t)
	case s == "3":
		fmt.Println("Введите ID")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID not found")
		}
		Done(id, *t)
	case s == "4":
		fmt.Println("Введите ID")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ID not found")
		}
		Delete(id, t)
	case s == "5":
		var str string
		fmt.Println("Введите строку")
		scanner.Scan()
		str = scanner.Text()
		TaskSearch(str, *t)
	case s == "6":
		TaskSort(*t)
	case s == "7":
		end = true
	}
	return end
}

func TaskSort(s []Task) { // Сортировка задач по ID
	sort.Slice(s, func(i, j int) bool {
		return s[i].ID < s[j].ID
	})
	fmt.Println(s)
}

func GetDeadlineStatus(tm time.Time) { //получение статуса дэдлайна
	tmNow := time.Now()
	switch {
	case tm.Before(tmNow):
		fmt.Println("Deadline status if overdue")
	case tm.After(tmNow):
		fmt.Println("Deadline status is upcoming")
	case tm.Equal(tmNow.Add(time.Hour)):
		fmt.Println("The deadline is in an hour")
	}
}

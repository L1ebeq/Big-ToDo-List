package models

import (
	"errors"
	"fmt"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func AddOneTaks(a Task, b *[]Task) {

	*b = append(*b, a)
}

func ListTasks(s []Task) {

	for _, v := range s {
		fmt.Println("ID:", v.ID)
		fmt.Println(v.Title)
		if v.Completed == true {
			fmt.Println("Статус: Выполнено")
		} else {
			fmt.Println("Статус: Не выполнено")
		}
	}

}

func Add(str string, s *[]Task) {
	var l int
	l = len(*s)
	var t = Task{
		ID:        l + 1,
		Title:     str,
		Completed: false,
	}
	*s = append(*s, t)
	fmt.Print(t)
}

func Done(id int, s []Task) (error, bool) {
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

func Delete(id int, s *[]Task) (error, bool) {
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

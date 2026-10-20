package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager/internal/models"
)

func SaveData(file *os.File, s []models.Task) { //Запись данных в файл
	os.Open(file.Name())
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Data marshal error")
	}
	file.Write(data)
	defer file.Close()
}

func ReadData(file *os.File) *[]models.Task { //Чтение данных из файла
	os.Open(file.Name())
	var data []byte
	var s *[]models.Task
	_, err := file.Read(data)
	if err != nil {
		fmt.Println("Read file error")
	}
	json.Unmarshal(data, s)
	return s
}

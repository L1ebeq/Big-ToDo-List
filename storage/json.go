package storage

import (
	"encoding/json"
	"io"
	"os"
	"task-manager/internal/models"
)

func SaveData(path string, s []models.Task) error { //Запись данных в файл
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadData(path string) ([]models.Task, error) { //Чтение данных из файла
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var s []models.Task
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

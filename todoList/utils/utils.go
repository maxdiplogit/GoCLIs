package utils

import (
	"errors"
	"os"
	"todo-list/todo"
)

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(os.ErrNotExist, err) {
			err = os.WriteFile(filePath, []byte("[]"), 0644)
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	return true, nil
}

func WriteToFile(filePath string, content []byte) (bool, error) {
	fileExists, err := fileExists(filePath)
	if err != nil {
		return false, err
	}

	if fileExists {
		err = os.WriteFile(filePath, content, 0644)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func ReadFromFile(filePath string) ([]byte, error) {
	fileExists, err := fileExists(filePath)
	if err != nil {
		return nil, err
	}

	if fileExists {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		return content, nil
	}

	return nil, nil
}

func QuickSort(todos []*todo.Todo, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := getPivotIndex(todos, low, high)
	QuickSort(todos, low, pivotIndex-1)
	QuickSort(todos, pivotIndex+1, high)
}

func getPivotIndex(todos []*todo.Todo, low, high int) int {
	pivot := todos[high].Id
	i := low - 1

	for j := low; j < high; j++ {
		if todos[j].Id <= pivot {
			i++
			todos[i], todos[j] = todos[j], todos[i]
		}
	}

	todos[i+1], todos[high] = todos[high], todos[i+1]
	return i + 1
}

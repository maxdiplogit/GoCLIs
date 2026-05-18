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
	todos[low], todos[(low+high)/2] = todos[(low+high)/2], todos[low]
	pivotElement := todos[low].Id

	i := low + 1
	j := high

	for i <= j {
		for i <= high && todos[i].Id < pivotElement {
			i++
		}

		for j >= low+1 && todos[j].Id > pivotElement {
			j--
		}

		if i <= j {
			todos[i], todos[j] = todos[j], todos[i]
			i++
			j--
		}
	}

	todos[low], todos[j] = todos[j], todos[low]
	return j
}

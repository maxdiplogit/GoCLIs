package main

import (
	// "encoding/json"
	// "errors"
	"encoding/json"
	// "errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"todo-list/todo"
	"todo-list/utils"
	// "path/filepath"
	// "todo-list/todo"
)

const ADD, DONE, UNDONE, LIST string = "add", "done", "undone", "list"
const MAX_TODOS = 10

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./todo-list <operation>\nOperation types: %s, %s, %s, %s\n%s compulsory operation flags: -d (Description)\n%s compulsory operation flags: -i (Index of Todo)", ADD, DONE, UNDONE, LIST, ADD, UNDONE)
		os.Exit(1)
	}

	userWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fileName := "todo.json"

	filePath := filepath.Join(userWorkingDir, fileName)

	fmt.Println("filePath: ", filePath)

	addCmd := flag.NewFlagSet(ADD, flag.ExitOnError)
	doneCmd := flag.NewFlagSet(DONE, flag.ExitOnError)
	undoneCmd := flag.NewFlagSet(UNDONE, flag.ExitOnError)

	description := addCmd.String("d", "test", "Todo Description")
	indexDone := doneCmd.Int("i", 0, "Index of Todo")
	indexUndone := undoneCmd.Int("i", 0, "Index of Todo")

	content, err := utils.ReadFromFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var todos []*todo.Todo = make([]*todo.Todo, 0, MAX_TODOS)
	err = json.Unmarshal(content, &todos)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case ADD:
		addCmd.Parse(os.Args[2:])

		if len(todos) == MAX_TODOS {
			fmt.Fprintf(os.Stderr, "Cannot add more todos, max %d allowed\n!", MAX_TODOS)
			os.Exit(1)
		}

		newTodo := todo.New((*todos[len(todos)-1]).Id+1, *description)
		todos = append(todos, newTodo)

		utils.QuickSort(todos, 0, len(todos)-1)

		todo.List(todos)
	case DONE:
		doneCmd.Parse(os.Args[2:])
		fmt.Println("Index: ", *indexDone)

		if *indexDone < 0 {
			fmt.Fprintln(os.Stderr, "index should be positive!")
			os.Exit(0)
		}

		if *indexDone >= len(todos) {
			fmt.Fprintln(os.Stderr, "index should be less than number of todos!")
			os.Exit(0)
		}

		todos[*indexDone].Edit(true)

		todo.List(todos)
	case UNDONE:
		undoneCmd.Parse(os.Args[2:])
		fmt.Println("Index: ", *indexUndone)

		if *indexUndone < 0 {
			fmt.Fprintln(os.Stderr, "index should be positive!")
			os.Exit(0)
		}

		if *indexUndone >= len(todos) {
			fmt.Fprintln(os.Stderr, "index should be less than number of todos!")
			os.Exit(0)
		}

		todos[*indexUndone].Edit(false)

		todo.List(todos)
	case LIST:
		todo.List(todos)
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}

	// Sort the todos on the basis of Id before writing them to file
	// Implement quick sort for this
	utils.QuickSort(todos, 0, len(todos)-1)

	finalContent, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	success, err := utils.WriteToFile(filePath, finalContent)

	if !success && err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

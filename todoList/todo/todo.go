package todo

import "fmt"

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"done"`
}

func New(id int, description string) *Todo {
	todo := &Todo{
		Id:          id,
		Description: description,
		Completed:   false,
	}

	return todo
}

func (todo *Todo) Edit(value bool) bool {
	(*todo).Completed = value

	return true
}

func List(todos []*Todo) {
	for _, todo := range todos {
		fmt.Println(*todo)
	}
}

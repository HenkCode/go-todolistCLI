package controller

import "fmt"

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo
var nextID = 1

func ViewTodos() {
	fmt.Println("Daftar Todo:")

	for _, todo := range todos {
		status := "Belum selesai"

		if todo.Done {
			status = "Selesai"
		}
		fmt.Printf("%d: %s [%s]\n", todo.ID, todo.Task, status)

	}
}

func AddTodos() {
	fmt.Println("Masukan List Tugas mu")

	var task string
	fmt.Scan(&task)

	todos = append(todos, Todo{
		ID : nextID,
		Task: task,
		Done: false,
	} )
}git 
package controller

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo
var nextID = 1

func saveTodos() {
	data, _ := json.Marshal(todos)
	os.WriteFile("todos.json", data, 0644)
}

func loadTodos() {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &todos)
	if len(todos) > 0 {
		nextID = todos[len(todos)-1].ID + 1
	}
}

func ViewTodos() {
	fmt.Println("Daftar Todo: ")

	for _, todo := range todos {
		status := "Belum selesai"

		if todo.Done {
			status = "Selesai"
		}
		fmt.Printf("%d: %s [%s]\n", todo.ID, todo.Task, status)

	}
}

func AddTodos() {
	fmt.Println("Masukan List Tugas mu: ")

	var task string
	fmt.Scan(&task)

	todos = append(todos, Todo{
		ID : nextID,
		Task: task,
		Done: false,
	})
	nextID++
	fmt.Println("Tugas Berhasil Ditambahkan")
}

func MarkDone() {
	fmt.Println("Masukan ID tugas yg selesai: ")

	var id int
	fmt.Scan(&id)

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			fmt.Println("Tugas ditandai selesai")
			return
		}
	}
	fmt.Println("ID tidak ditemukan")
}

func DeleteTodos() {
	fmt.Println("Masukan ID yg ingin dihapus: ")
	var id int
	fmt.Scan(&id)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Tugas dihapus")
			return
		}
	}
	fmt.Println("ID tidak ditemukan")
}
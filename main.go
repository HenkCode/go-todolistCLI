package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo
var nextID = 1

// func saveTodos() {
// 	data, _ := json.Marshal(todos)
// 	os.WriteFile("todos.json", data, 0644)
// }

// func loadTodos() {
// 	data, err := os.ReadFile("todos.json")
// 	if err != nil {
// 		return
// 	}
// 	json.Unmarshal(data, &todos)
// 	if len(todos) > 0 {
// 		nextID = todos[len(todos)-1].ID + 1
// 	}
// }

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

func AddTodos(tasks string) {
	todos = append(todos, Todo{
		ID : nextID,
		Task: tasks,
		Done: false,
	})
	nextID++
	fmt.Println("Tugas Berhasil Ditambahkan")
}

func MarkDone(i int) {
	if i >= 1 && i <= len(todos) {
		todos[i-1].Done = true
		fmt.Println("Tugas selesai")
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}

func DeleteTodos(i int) {
	if i >= 1 && i <= len(todos) {
		todos = append(todos[:i-1], todos[i:]...)
		fmt.Println("berhasil dihapus")
	} else {
		fmt.Println("ID tidak ditemukan")
	}
}

func main() {
	fmt.Println("Todo List CLI")
	fmt.Println("--------------")
	fmt.Println("1. Lihat Todo")
	fmt.Println("2. Tambah Todo")
	fmt.Println("3. Tandai Selesai")
	fmt.Println("4. Hapus Todo")
	fmt.Println("5. Keluar")

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pilih opsi: ")

		scanner.Scan()
		input := scanner.Text()

		choice, _ := strconv.Atoi(input)

		switch choice {
		case 1:
			ViewTodos()
		case 2:
			fmt.Println("Masukan Tugas mu: ")
			scanner.Scan()
			taskInput := scanner.Text()
			AddTodos(taskInput)
		case 3:
			fmt.Println("Masukan ID tugas yg selesai: ")
			scanner.Scan()
			iinput, _ := strconv.Atoi(scanner.Text())
			MarkDone(iinput)
		case 4:
			fmt.Println("Masukan ID yg ingin dihapus: ")
			scanner.Scan()
			iinput, _ := strconv.Atoi(scanner.Text())
			DeleteTodos(iinput)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Opsi Tidak Ada")
			
		}
		
	}
}
package main

import (
	"fmt"
	"os"

	"github.com/HenkCode/go-todolistCLI/controllers"
)



func main() {
	for {
		fmt.Println("Todo List CLI")
		fmt.Println("--------------")
        fmt.Println("1. Lihat Todo")
        fmt.Println("2. Tambah Todo")
        fmt.Println("3. Tandai Selesai")
        fmt.Println("4. Hapus Todo")
        fmt.Println("5. Keluar")
        fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			controller.ViewTodos()
		case 2:
			controller.AddTodos()
		case 3:
			controller.MarkDone()
		case 4:
			controller.DeleteTodos()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Opsi Tidak Ada")
			
		}
		
	}
}
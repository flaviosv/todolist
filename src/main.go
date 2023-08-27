package main

import (
	"fmt"
	"os"
	"todolist/view"
)
import "todolist/model"

func main() {
	view.PrintGreeting()

	model.List()

	for {
		option := view.Menu()

		switch option {
		case 1:
			model.PrepareAddTodo()
		case 2:
			model.PrepareEditTodo()
		case 5:
			model.List()
		case 99:
			fmt.Println("Thanks for using this incredible Todolist!")
			os.Exit(1)
		default:
			fmt.Println("Option doesn't exist")
		}
	}

}

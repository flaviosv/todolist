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

		if option == 99 {
			fmt.Println("Thanks for using this incredible Todolist!")
			os.Exit(1)
		}

		switch option {
		case 1:
			model.PrepareAddTodo()
		}
	}

}

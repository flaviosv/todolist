package view

import "fmt"

func PrintGreeting() {
	fmt.Println("Welcome to your Todolist!")
}

func Menu() int {
	fmt.Println("\nWhat do you want to do? Choose an option")
	fmt.Println("1 - Add (Todo)?")
	fmt.Println("2 - Edit (Todo)?")
	fmt.Println("3 - Mark as done/undone (Todo)?")
	fmt.Println("4 - Delete (Todo)?")
	fmt.Println("99 - Exit program?")

	var option int
	fmt.Scanln(&option)

	return option
}

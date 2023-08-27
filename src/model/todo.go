package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Filename       = "todo.txt"
	FileHeader     = "title,status,createdAt,doneAt"
	IndexTitle     = 0
	IndexStatus    = 1
	IndexCreatedAt = 2
	IndexDoneAt    = 3
)

type todo struct {
	title     string
	status    int
	createdAt time.Time
	doneAt    time.Time
}

func (t todo) toString() string {

	return fmt.Sprintf("%s,%s,%s,%s", t.title, strconv.Itoa(t.status), t.createdAt.Format(time.RFC3339), t.doneAt.Format(time.RFC3339))
}

func List() {
	list := getList()

	if len(list) <= 0 {
		fmt.Println("There isn't any todos in our list")
	}

	fmt.Println("\nYour todolist")
	for i, t := range list {
		status := "Done"
		if t.status == 0 {
			status = "Not done"
		}

		doneAt := ""
		if t.status == 1 {
			doneAt = t.doneAt.Format("2006-01-02 15:04")
		}

		fmt.Println(i+1, "-", t.title, "-", status, "-", doneAt)
	}
}

func getList() []todo {
	file, err := os.ReadFile(Filename)

	if err != nil {
		createBlankfile()

		return []todo{}
	}

	content := strings.Split(string(file), "\n")
	todos := []todo{}
	for i, line := range content {
		if i == 0 {
			continue
		}

		if line == "" {
			continue
		}

		lineSlice := strings.Split(line, ",")
		status, _ := strconv.Atoi(lineSlice[IndexStatus])
		createdAt, _ := time.Parse(time.RFC3339, lineSlice[IndexCreatedAt])
		doneAt, _ := time.Parse(time.RFC3339, lineSlice[IndexDoneAt])
		todos = append(todos, todo{
			title:     lineSlice[IndexTitle],
			status:    status,
			createdAt: createdAt,
			doneAt:    doneAt,
		})
	}

	return todos
}

func createBlankfile() os.File {
	os.Remove(Filename)

	os.WriteFile(Filename, []byte(FileHeader), 0666)

	f, _ := os.OpenFile(Filename, os.O_WRONLY, 0666)

	return *f
}

func PrepareAddTodo() {
	fmt.Println("What is it?")

	var title string
	fmt.Scanln(&title)

	err := appendTodo(title)

	evaluateOperationResult(err)
}

func evaluateOperationResult(err error) {
	if err != nil {
		fmt.Println("It presented an error", err)

		return
	}

	fmt.Println("Done!")
}

func appendTodo(title string) error {
	f, _ := os.OpenFile(Filename, os.O_APPEND|os.O_WRONLY, 0666)

	t := todo{
		title:     title,
		status:    0,
		createdAt: time.Now(),
		doneAt:    time.Now(),
	}

	defer f.Close()
	f.WriteString("\n")
	_, err := f.WriteString(t.toString())

	return err
}

func PrepareEditTodo() {
	todoCode := scanTodoCode()

	todoExist := isTodoCodeExist(todoCode)
	if !todoExist {
		return
	}

	todoList := getList()
	todo := todoList[todoCode-1]
	fmt.Println("What should be the new name of the Todo '", todo.title, "'")

	var title string
	fmt.Scanln(&title)

	err := edit(title, todoCode-1)

	evaluateOperationResult(err)
}

func scanTodoCode() int {
	fmt.Println("What is the todo number you want to work with?")

	var todoCode int
	fmt.Scanln(&todoCode)

	return todoCode
}

func isTodoCodeExist(todoCode int) bool {
	todoList := getList()
	if todoCode > len(todoList) {
		fmt.Println("Todo doesn't exist")

		return false
	}

	return true
}

func edit(title string, index int) error {
	todoList := getList()
	todoList[index].title = title

	return saveCompleteList(todoList)
}

func saveCompleteList(todoList []todo) error {
	f := createBlankfile()

	strList := FileHeader + "\n"
	for _, todo := range todoList {
		strList = strList + todo.toString() + "\n"
	}

	_, err := f.WriteString(strList)

	return err
}

func PrepareMarkDoneUndone() {
	todoCode := scanTodoCode()

	todoExist := isTodoCodeExist(todoCode)
	if !todoExist {
		return
	}

	todoList := getList()
	todo := todoList[todoCode-1]
	action := "Done"
	if todo.status == 1 {
		action = "Not done"
	}
	fmt.Println("Do you want to mark '", todo.title, "' as", action, "? (y/n)")

	conf := getConfirmationInput()
	if !conf {
		return
	}

	newStatus := 0
	if todo.status == 0 {
		newStatus = 1
	}
	err := changeStatus(todoCode-1, newStatus)

	evaluateOperationResult(err)
}

func getConfirmationInput() bool {
	var conf string
	fmt.Scanln(&conf)

	options := map[string]int{
		"y": 0,
		"n": 0,
	}
	if _, ok := options[conf]; !ok {
		fmt.Println("Options invalid!")

		return false
	}

	if conf == "n" {
		fmt.Println("Ok, getting back to the menu!")

		return false
	}

	return true
}

func changeStatus(index int, newStatus int) error {
	todoList := getList()
	todoList[index].status = newStatus
	todoList[index].doneAt = time.Now()

	return saveCompleteList(todoList)
}

func PrepareDelete() {
	todoCode := scanTodoCode()

	todoExist := isTodoCodeExist(todoCode)
	if !todoExist {
		return
	}

	todoList := getList()
	todo := todoList[todoCode-1]
	fmt.Println("Do you really want to remove the item '", todo.title, "'? (y/n)")

	conf := getConfirmationInput()
	if !conf {
		return
	}

	err := deleteTodo(todoCode - 1)

	evaluateOperationResult(err)
}

func deleteTodo(index int) error {
	todoList := getList()
	newTodoList := append(todoList[:index], todoList[index+1:]...)

	return saveCompleteList(newTodoList)
}

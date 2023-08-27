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
	fmt.Println("What is the todo number you want to edit?")

	var todoCode int
	fmt.Scanln(&todoCode)

	todoList := getList()
	if todoCode > len(todoList) {
		fmt.Println("Todo doesn't exist")

		return
	}

	fmt.Println("What should be the new name of the Todo '", todoList[todoCode-1].title, "'")

	var title string
	fmt.Scanln(&title)

	err := edit(title, todoCode-1)

	if err != nil {
		fmt.Println("It presented an error", err)

		return
	}

	fmt.Println("Done!")
}

func edit(title string, index int) error {
	todoList := getList()
	for i, _ := range todoList {
		if i == index {
			todoList[i].title = title
		}
	}

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

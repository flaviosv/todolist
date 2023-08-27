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

func List() {
	list := getList()

	if len(list) <= 0 {
		fmt.Println("There isn't any todos in our list")
	}

	fmt.Println("Your todolist")
	for i, t := range list {
		var status = "Done"
		if t.status == 0 {
			status = "Not done"
		}

		fmt.Println(i+1, "-", t.title, "-", status, "-", t.doneAt.Format("2006-01-02 15:04"))
	}
}

func getList() []todo {
	file, err := os.ReadFile(Filename)

	if err != nil {
		os.WriteFile(Filename, []byte(FileHeader), 0666)

		return []todo{}
	}

	content := strings.Split(string(file), "\n")
	todos := []todo{}
	for i, line := range content {
		if i == 0 {
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

package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	menu "github.com/benjaNygit/cli-crud/go/src/menu"
	task "github.com/benjaNygit/cli-crud/go/src/tasks"
)

func main() {
	file := FileMain()
	defer file.Close()
	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// check al tamaño del archivo
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tasks = []task.Task{}
	}

	menu.Menu(tasks, file)
}

func FileMain() *os.File {
	file, err := os.OpenFile("data/tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

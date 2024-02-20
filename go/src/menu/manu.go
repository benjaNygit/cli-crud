package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	task "github.com/benjaNygit/cli-crud/go/src/models"
)

func Menu(tasks []task.Task, file *os.File) {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Nombre de la tarea a guardar:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = append(tasks, task.AddTask(name))
		task.SaveTasks(tasks, file)
	case "list":
		task.ListTasks(tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Especifica el indice de la tarea")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
			return
		}

		tasks = task.CompletedTask(tasks, id)
		task.SaveTasks(tasks, file)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Especifica el indice de la tarea")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(tasks, file)

	default:
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println("Uso: [add|list|complete|delete]")
}

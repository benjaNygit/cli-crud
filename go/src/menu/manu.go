package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	task "github.com/benjaNygit/cli-crud/go/src/tasks"
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
		return
	case "delete":
		return

	default:
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println("Uso: [add|list|complete|delete]")
}

package menu

import (
	"fmt"
	"os"

	task "github.com/benjaNygit/cli-crud/go/src/tasks"
)

func Menu(tasks []task.Task, file *os.File) {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

}

func PrintUsage() {
	fmt.Println("Uso: [add|list|complete|delete]")
}

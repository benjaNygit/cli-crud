package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Completed    bool      `json:"completed"`
	Created_at   time.Time `json:"created_at"`
	Completed_at time.Time `json:"completed_at"`
}

func SaveTasks(tasks []Task, file *os.File) {
	// convertir a json
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	// mover al comienzo del archivo
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// limpiar archivo
	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	// volver a escribir el archivo
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}

	// asegurarse que se escribió el contenido
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func AddTask(name string) Task {
	task := Task{
		Id:           uuid.New(),
		Name:         name,
		Completed:    false,
		Created_at:   time.Now(),
		Completed_at: time.Time{},
	}

	return task
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%d - %t %s \n", i+1, task.Completed, task.Name)
	}
}

func CompletedTask(tasks []Task, id int) []Task {
	for i := range tasks {
		if i == id-1 {
			tasks[i].Completed = true
			break
		}
	}

	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i := range tasks {
		if i == id-1 {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}

	return tasks
}

package main

import (
	"log"
	"os"
)

func main() {
	// file := FileMain()
}

func FileMain() *os.File {
	file, err := os.OpenFile("data/tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return file
}

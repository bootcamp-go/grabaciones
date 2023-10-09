package main

import (
	"fmt"
	"time"
)

func work(task string) {
	time.Sleep(2 * time.Second)
	fmt.Println(task, "done")
}

func main() {
	tasks := []string{"task1", "task2", "task3"}

	for _, task := range tasks {
		go work(task)
	}

	time.Sleep(3 * time.Second)
}
package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// import (
// 	"github.com/satori/go.uuid"
// )

// Task struct
type Task struct {
	id   string
	name string
	desc string
}

// State => TODO: Replace the slice with an actual database
var tasks []Task

// Add task => TODO: Add input validation
func addTask(name, desc string) {
	taskID := uuid.Must(uuid.NewV4()).String()
	tasks = append(tasks, Task{taskID, name, desc})
}

// Show tasks
func showTasks() {
	for _, task := range tasks {
		fmt.Printf("\nTask ID: %s\nTask Name: %s\nTask Description: %s\n", task.id, task.name, task.desc)
	}
}

// Update task
func updateTask(task *Task, name, desc string) {
	task.name = name
	task.desc = desc
}

// Delete task => TODO: Learn how the **** does this function even work - Thanks StackOverflow
func deleteTask(id string) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}

func main() {
	addTask("Test One", "Anna One")
	addTask("Test Two", "Anna Two")
	addTask("Test Three", "Anna one, two, three, four.")
	updateTask(&tasks[2], "Edited Task", "damn son where did u find dis")
	deleteTask(tasks[1].id)
	showTasks()
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var taskItems = []string{
	"Watch Go crash course",
	"Watch Nana's Golang Full Course",
	"Reward myself with a donut",
}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)
	http.HandleFunc("/update-task", updateTask)
	http.HandleFunc("/delete-task", deleteTask)

	http.ListenAndServe(":8080", nil)
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user. Welcome to our Todolist App!")
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for i, task := range taskItems {
		fmt.Fprintf(w, "%d: %s\n", i, task)
	}
}

func addTask(w http.ResponseWriter, r *http.Request) {
	task := r.URL.Query().Get("task")

	if task == "" {
		fmt.Fprintln(w, "Please provide task")
		return
	}

	taskItems = append(taskItems, task)
	fmt.Fprintln(w, "Task added!")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")
	newTask := r.URL.Query().Get("task")

	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(taskItems) {
		fmt.Fprintln(w, "Invalid index")
		return
	}

	taskItems[index] = newTask
	fmt.Fprintln(w, "Task updated!")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")

	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(taskItems) {
		fmt.Fprintln(w, "Invalid index")
		return
	}

	taskItems = append(taskItems[:index], taskItems[index+1:]...)
	fmt.Fprintln(w, "Task deleted!")
}

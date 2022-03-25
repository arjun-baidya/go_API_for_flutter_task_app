package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       "2020-01-01"}
	tasks = append(tasks, task)
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       "2020-01-02"}
	tasks = append(tasks, task2)
	fmt.Println("all task", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	fmt.Println("I am gettasks page")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			flag = true
			json.NewEncoder(w).Encode(tasks[i])
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"error": "Task not found"})
	}

	fmt.Println("I am gettask page")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am createtask page")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am updatetask page")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am deletetask page")
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	allTasks()
	fmt.Println("Hello, World!")
	handleRoutes()
}

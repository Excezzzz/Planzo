package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type task struct {
	text   string
	id     int
	status bool
}

var tasks []task

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/list", tasklistHandler)

	log.Println("server start")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		gettask(w, r)
	case http.MethodPost:
		posttask(w, r)
	default:
		http.Error(w, "ivalid method", http.StatusMethodNotAllowed)
	}
}

func gettask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
	fmt.Println("get tasks", tasks)
}

func posttask(w http.ResponseWriter, r *http.Request) {
	var tasks task
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("post new task", tasks)
}

func tasklistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("web serner works correctly")
}

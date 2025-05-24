package main

import (
	"encoding/json"
	"net/http"
)

type task struct {
	task   string
	id     int
	status bool
}

var tasks []task

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/list", tasklistHandler)
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
}

func posttask(w http.ResponseWriter, r *http.Request) {
	var t task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tasks = append(tasks, t)
	w.WriteHeader(http.StatusCreated)
}

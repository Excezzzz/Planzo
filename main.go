package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

type task struct {
	Text   string `json:"text"`
	ID     int    `json:"id"`
	Status bool   `json:"status"`
}

var tasks []task
var nextID int

func main() {
	http.HandleFunc("/addtask", addtaskHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/deletetask", deletetaskHandler)

	log.Println("server start")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func addtaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addtask(w, r)
	default:
		http.Error(w, "ivalid method", http.StatusMethodNotAllowed)
		fmt.Println("addtask invalid method")
	}
}

func addtask(w http.ResponseWriter, r *http.Request) {
	var t task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = nextID
	nextID++

	tasks = append(tasks, t) // Добавляем в глобальный слайс

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "task added")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks)
		fmt.Println("tasks list send", tasks)
	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		fmt.Println("list invalid method")
	}

}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		changestatus(w, r)
	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

func changestatus(w http.ResponseWriter, r *http.Request) {
	type update struct {
		ID     int
		Status bool
	}

	var upd update
	err := json.NewDecoder(r.Body).Decode(&upd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("bad request")
	}

	found := false
	for i, t := range tasks {
		if t.ID == upd.ID {
			tasks[i].Status = upd.Status
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "task not found. check ID", http.StatusNotFound)
		fmt.Println("task not found. check ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "status updated")
}

func deletetaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		deletetask(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func deletetask(w http.ResponseWriter, r *http.Request) {
	type deletetask struct {
		ID int
	}

	var del deletetask
	err := json.NewDecoder(r.Body).Decode(&del)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	founddel := false
	for i, t := range tasks {
		if t.ID == del.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if !founddel {
		http.Error(w, "task not found. check ID", http.StatusNotFound)
		fmt.Println("task not found. check ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "task deleted")
}

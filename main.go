package main

import (
	"fmt"
	"net/http"
)

type task struct {
	id int
	text string
	done bool
}

func main() {
	// Регистрируем маршрут
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			task := r.FormValue("task")
			fmt.Fprintf(w, "Ты добавил задачу: %s", task)
		}
	})

	// Запускаем сервер на 8080 порту
	fmt.Println("Сервер на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

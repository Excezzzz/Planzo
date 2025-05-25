package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

type task struct {
	Text   string
	ID     int
	Status bool
}

var tasks []task
var nextID int

func main() {
	fmt.Println("Server start on loclahost:8080")

	e := echo.New()

	e.POST("/addtask", addtaskHandler)
	e.GET("/list", listHandler)

	e.Start(":8080")
}

func addtaskHandler(c echo.Context) error {
	var t task
	if err := c.Bind(&t); err != nil {
		log.Println("Bad request")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	t.ID = nextID
	nextID++
	t.ID = len(tasks) + 1

	tasks = append(tasks, t)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	log.Println("Task created")
	return c.String(http.StatusCreated, "Task created")
}

func listHandler(c echo.Context) error {
	log.Println("Tasks list send")
	return c.JSON(http.StatusOK, tasks)
}

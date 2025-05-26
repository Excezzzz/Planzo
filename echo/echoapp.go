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
	e.PATCH("/status", statusHandler)
	e.DELETE("/deletetask", deletetaskHandler)

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

func statusHandler(c echo.Context) error {
	type update struct {
		ID     int
		Status bool
	}

	var upd update
	if err := c.Bind(&upd); err != nil {
		log.Println("Bad request")
		return c.JSON(http.StatusBadRequest, err.Error())
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
		log.Println("Task not found. Check ID")
		return c.JSON(http.StatusNotFound, "Task not found. Check ID")
	}

	log.Println("Status updated")
	return c.JSON(http.StatusOK, "Status updated")
}

func deletetaskHandler(c echo.Context) error {
	type deletetask struct {
		ID int
	}

	var del deletetask
	if err := c.Bind(&del); err != nil {
		log.Println("Bad request")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	founddel := false
	for i, t := range tasks {
		if t.ID == del.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			reorderIDs()
			log.Println("Task deleted")
			return c.JSON(http.StatusOK, "Task deleted")
		}
	}

	if !founddel {
		log.Println("Task not found. Check ID")
		return c.JSON(http.StatusNotFound, "Task not found. Check ID")
	}

	return c.JSON(http.StatusOK, "Task deleted")
}

func reorderIDs() {
	for i := range tasks {
		tasks[i].ID = i + 1
	}
}

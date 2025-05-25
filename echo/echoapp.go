package main

import (
	"fmt"
	"net/http"

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

	e.GET("/addtask", func(c echo.Context) error {
		addtaskHandler()
		return c.String(http.StatusOK, "/addtask")
	})

	e.Start(":8080")
}

func addtaskHandler() {

}

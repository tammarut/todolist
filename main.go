package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tammarut/todolist-recap/handler"
)

type List struct {
	Topic string `json: "topic"`
}

func main() {
	// Echo instance
	e := echo.New()

	// MiddlewareW
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n", //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))
	e.Use(middleware.Recover())

	// Intial Handler
	h := handler.Handler{}

	// Routes
	e.GET("/", h.Hello)
	e.POST("/todos", saveList)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func saveList(c echo.Context) error {
	list := new(List)
	if err := c.Bind(list); err != nil {
		log.Println("Error: from saveList", err)
		c.String(http.StatusInternalServerError, "Error: from c.Bind func saveList")
	}
	fmt.Printf("%v#\n", list)
	return c.JSON(http.StatusCreated, list)
}

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
	ID    string `json:"id"`
	Title string `json:"title"`
}

var lists []List //=> global list

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
	e.GET("/", h.Hello)              //=> Hello world
	e.GET("/todos/:id", getListById) //=> get list by id
	e.GET("/todos", getAllLists)     //=> get all lists
	e.POST("/todos", saveList)       //=> post list from body

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func saveList(c echo.Context) error { //=> post only 1
	var l List
	if err := c.Bind(&l); err != nil {
		log.Println("Error: from saveList", err)
	}

	lists = append(lists, l)
	fmt.Printf("%#v\n", l)
	return c.JSON(http.StatusCreated, l)
}

func getAllLists(c echo.Context) error { //=> get all lists: OK
	return c.JSON(http.StatusOK, &lists)
}

func getListById(c echo.Context) error { //=> get 1 list by id
	id := c.Param("id")
	for i, v := range lists {
		if v.ID == id {
			return c.JSON(http.StatusOK, lists[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Not found this id")
}

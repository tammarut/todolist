package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tammarut/todolist/model"
)

var lists []model.List

func main() {
	// Echo instance
	e := echo.New()

	// MiddlewareW
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n", //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Hello)                  //=> Hello world
	e.GET("/todos/:id", getListByID)   //=> get list by id
	e.GET("/todos", getAllLists)       //=> get all lists
	e.POST("/todos", saveList)         //=> post list from body
	e.DELETE("/todos/:id", deleteByID) //=>delete a list by id
	e.PATCH("/todos/:id", updateByID)  //=>update a list bu id

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello todolist")
}
func saveList(c echo.Context) error { //=> post only 1
	var l model.List
	if err := c.Bind(&l); err != nil {
		log.Println("Error: from saveList", err)
	}
	lists = append(lists, l)
	fmt.Printf("%#v\n", l)
	return c.JSON(http.StatusCreated, lists)
}

func getAllLists(c echo.Context) error { //=> get all lists: OK
	return c.JSON(http.StatusOK, &lists)
}

func getListByID(c echo.Context) error { //=> get 1 list by id
	id := c.Param("id")
	for i := range lists {
		if lists[i].ID == id {
			return c.JSON(http.StatusOK, lists[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Not found this id")
}
func deleteByID(c echo.Context) error {
	id := c.Param("id")

	filterLists := []model.List{}
	for _, item := range lists {
		if item.ID != id {
			filterLists = append(filterLists, item)
		}
	}

	lists = filterLists
	return c.JSON(http.StatusOK, lists)
}

func updateByID(c echo.Context) error {
	newtitle := new(model.List)
	if err := c.Bind(newtitle); err != nil {
		log.Println("Error: from updateByID", err)
	}

	id := c.Param("id")
	for i := range lists {
		if lists[i].ID == id {
			fmt.Printf("Before: %+v\n", lists[i])
			lists[i].Title = newtitle.Title
			fmt.Printf("After: %+v\n", lists[i])
			return c.JSON(http.StatusOK, lists[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Not found this ID!",
	})
}

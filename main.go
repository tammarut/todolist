package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tammarut/todolist/handler"
)

var (
	logFormat = `[${time_rfc3339}]  status=${status}  ${method} ${host}${path}  ${latency_human} ${latency}` + "\n"
)

func main() {
	// Echo instance
	e := echo.New()

	// MiddlewareW
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ //=> Custom log
		Format: logFormat, //=> [2019-02-11T23:56:04+07:00]  status=200  GET localhost:1323/admin/main  0s
	}))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handler.Hello)                  //=> Hello world
	e.GET("/todos/:id", handler.GetListByID)   //=> get list by id
	e.GET("/todos", handler.GetAllLists)       //=> get all lists
	e.POST("/todos", handler.SaveList)         //=> post list from body
	e.DELETE("/todos/:id", handler.DeleteByID) //=>delete a list by id
	e.PATCH("/todos/:id", handler.UpdateByID)  //=>update a list bu id

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

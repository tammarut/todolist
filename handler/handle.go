package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tammarut/todolist/model"
)

var lists []model.List

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello todolist")
}
func SaveList(c echo.Context) error { //=> post only 1
	var l model.List
	if err := c.Bind(&l); err != nil {
		log.Println("Error: from saveList", err)
	}
	lists = append(lists, l)
	fmt.Printf("%+v\n", l)
	return c.JSON(http.StatusCreated, "We've created your list!")
}

func GetAllLists(c echo.Context) error { //=> get all lists: OK
	return c.JSON(http.StatusOK, &lists)
}

func GetListByID(c echo.Context) error { //=> get 1 list by id
	id := c.Param("id")
	for i := range lists {
		if lists[i].ID == id {
			return c.JSON(http.StatusOK, lists[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Not found this ID!",
	})
}
func DeleteByID(c echo.Context) error {
	id := c.Param("id")

	filterLists := []model.List{}
	for _, item := range lists {
		if item.ID != id {
			filterLists = append(filterLists, item)
		}
	}
	if len(lists) != len(filterLists) {
		lists = filterLists
		fmt.Printf("%+v\n", lists)
		return c.JSON(http.StatusOK, lists)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Not found this ID!",
		})
	}

}

func UpdateByID(c echo.Context) error {
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
			return c.JSON(http.StatusOK, lists)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Not found this ID!",
	})
}

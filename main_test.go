package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/tammarut/todolist/model"
)

var (
	//=>Mock-up AllLists because it runs on memory.
	mockup = []model.List{
	{
      ID:"1",
      Title:"excercise",
    },   
    {
	  ID:"2",
      Title:"play game",
   	},
   	{
      ID :"3",
      Title:"sleeping",
   	},
}
	//=>Expected
	allLists = `[{"id":"1","title":"excercise"},{"id":"2","title":"play game"},{"id":"3","title":"sleeping"}]`
)

func TestHelloShouldReturnHellotodolist(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil) //=>request GET "/"; no payload!
	rec := httptest.NewRecorder()                        //=>initial response
	c := e.NewContext(req, rec)

	//.Assertions
	if assert.NoError(t, Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)             //=>compare StatusCode; 'Want' vs 'rec'(response)
		assert.Equal(t, "Hello todolist", rec.Body.String()) //=>compare 'Want' vs 'mockDB'(resonse)
	}
}

func TestGetAllListsShouldReturnAllOfTodoLists(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos") //=>set path(url)
	lists = mockup

	//.Assertions
	if assert.NoError(t, getAllLists(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, allLists+"\n", rec.Body.String())
	}
}

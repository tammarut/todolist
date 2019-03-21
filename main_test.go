package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/tammarut/todolist/model"
)

var (
	//=>Mock-up Lists because it runs on memory.
	mockAllLists = []model.List{
		{
			ID:    "1",
			Title: "excercise",
		},
		{
			ID:    "2",
			Title: "play game",
		},
		{
			ID:    "3",
			Title: "sleeping",
		}}

	//=>Expected
	allLists = `[{"id":"1","title":"excercise"},{"id":"2","title":"play game"},{"id":"3","title":"sleeping"}]`
	AList    = `{"id":"2","title":"play game"}`
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
	c.SetPath("/todos")  //=>set path(url)
	lists = mockAllLists //=>dump(mock) new lists

	//.Assertions
	if assert.NoError(t, getAllLists(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, allLists+"\n", rec.Body.String())
	}
}
func TestGetListByIDWhenGotParamShouldReturnOneList(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, getListByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, AList+"\n", rec.Body.String())
	}
}

func TestSaveListWhenGotBodyRequest(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(AList))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//.Assertions
	if assert.NoError(t, saveList(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "\"We've created your list!\"\n", rec.Body.String())
	}
}

func TestGetAllListsShouldReturnAllOfTodoLists


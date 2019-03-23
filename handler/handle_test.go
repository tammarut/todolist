package handler

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
	mockUpdate = `{"title":"coding todolist"}`

	//=>Expected
	allLists         = `[{"id":"1","title":"excercise"},{"id":"2","title":"play game"},{"id":"3","title":"sleeping"}]`
	AList            = `{"id":"2","title":"play game"}`
	ListsWithDeleted = `[{"id":"1","title":"excercise"},{"id":"3","title":"sleeping"}]`
	ListsWithUpdated = `[{"id":"1","title":"excercise"},{"id":"2","title":"coding todolist"},{"id":"3","title":"sleeping"}]`
	NotFoundID       = `{"message":"Not found this ID!"}`
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
	if assert.NoError(t, GetAllLists(c)) {
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
	if assert.NoError(t, GetListByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, AList+"\n", rec.Body.String())
	}
}
func TestGetListByIDWhenGotWrongParamShouldReturnNotFoundID(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("10")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, GetListByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, NotFoundID+"\n", rec.Body.String())
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
	if assert.NoError(t, SaveList(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "\"We've created your list!\"\n", rec.Body.String())
	}
}

func TestDeleteByIDWhenGotParamShouldReturnAllListsWithoutIt(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, DeleteByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ListsWithDeleted+"\n", rec.Body.String())
	}
}
func TestDeleteByIDWHenGotWrongParamShouldReturnNotfoundID(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("10")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, DeleteByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, NotFoundID+"\n", rec.Body.String())
	}
}

func TestUpdateByIDWhenGotParamAndBodyShouldReturnUpdatedList(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(mockUpdate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, UpdateByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ListsWithUpdated+"\n", rec.Body.String())
	}
}

func UpdateByIDWhenGotWrongParamShouldReturnNotfoundID(t *testing.T) {
	//.Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(mockUpdate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")
	lists = mockAllLists

	//.Assertions
	if assert.NoError(t, UpdateByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, NotFoundID, rec.Body.String())
	}
}

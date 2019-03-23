package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
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

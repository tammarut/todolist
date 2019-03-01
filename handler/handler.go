package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	List struct {
		Topic string `json:"topic"`
	}

	Handler struct {
		DB map[string]*List
	}
)

//Handler
func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

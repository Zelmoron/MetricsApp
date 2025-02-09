package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoints struct {
}

func New() *Endpoints {
	return &Endpoints{}
}

func (e *Endpoints) Render(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", nil)
}

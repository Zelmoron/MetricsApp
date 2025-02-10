package endpoints

import (
	"MetricsApp/internal/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
		GetData(requests.FormData)
	}
	Endpoints struct {
		services Services
	}
)

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}

func (e *Endpoints) Render(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", nil)
}

func (e *Endpoints) GetData(c echo.Context) error {

	formData := requests.FormData{}
	if err := c.Bind(&formData); err != nil {
		c.Logger().Errorf("Ошибка привязки данных: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Неверные данные формы",
		})
	}

	c.Logger().Infof("Полученные данные: %+v", formData)
	e.services.GetData(formData)
	return nil
}

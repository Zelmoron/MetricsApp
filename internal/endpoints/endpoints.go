package endpoints

import (
	"MetricsApp/internal/requests"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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
	log.Info().Msgf("%s", formData.JSONData)
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

type Post struct {
	Name string `json:"name"`
}

func (e *Endpoints) ParallelPost(c echo.Context) error {
	var post Post
	if err := c.Bind(&post); err != nil {
		c.Logger().Errorf("Ошибка привязки данных: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Неверные данные формы",
		})
	}
	fmt.Println(post)
	return c.JSON(200, post)
}

package app

import (
	"MetricsApp/internal/endpoints"
	"MetricsApp/internal/services"
	"MetricsApp/internal/utils"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type App struct {
	app       *echo.Echo
	endpoints *endpoints.Endpoints
	services  *services.Services
	utils     *utils.Utils
}

func New() *App {

	a := &App{}
	a.app = echo.New()
	a.utils = utils.New()
	a.services = services.New()
	a.endpoints = endpoints.New(a.services)
	a.controllers()

	return a
}

func (a *App) controllers() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	a.app.Renderer = t
	a.app.Debug = true
	a.app.Static("/static", "static")
	a.app.Use(middleware.Recover())
	a.app.GET("/", a.endpoints.Render)
	a.app.POST("/get-data", a.endpoints.GetData)
	a.app.POST("/post", a.endpoints.ParallelPost)
}

func (a *App) Run() {
	url := "http://localhost:8080"

	a.utils.OpenBrowser(url)
	log.Info().Msg("Приложение запущено")
	a.app.Start(":8080")

}

package app

import (
	"MetricsApp/internal/endpoints"
	"io"
	"os/exec"
	"runtime"
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
}

func New() *App {

	a := &App{}
	a.app = echo.New()
	a.endpoints = endpoints.New()
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
	a.app.Use(middleware.Logger(), middleware.Recover())
	a.app.GET("/", a.endpoints.Render)
}

func (a *App) Run() {
	url := "http://localhost:8080"

	a.openBrowser(url)
	log.Info().Msg("Приложение запущено")
	a.app.Start(":8080")

}

func (a *App) openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		log.Warn().Msg("Не удалось определить ОС для открытия браузера")
		return
	}

	if err != nil {
		log.Error().Err(err).Msg("Не удалось открыть браузер")
	}
}

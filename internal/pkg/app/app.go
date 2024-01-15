package app

import (
	"fmt"
	"log"

	"github.com/diamondhulk625/middleware/internal/app/data"
	"github.com/diamondhulk625/middleware/internal/app/mw"
	"github.com/diamondhulk625/middleware/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	d    *data.Data
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.d = data.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.StepCheck)

	a.echo.GET("/status", a.d.Status)
	return a, nil

}

func (a *App) Run() error {
	fmt.Println("Server runing")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

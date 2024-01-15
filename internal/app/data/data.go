package data

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}
type Data struct {
	s Service
}

func New(s Service) *Data {
	return &Data{
		s: s,
	}
}

func (d *Data) Status(ctx echo.Context) error {
	t := d.s.DaysLeft()
	s := fmt.Sprintf("Left days until 01.01.2026: %d", t)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

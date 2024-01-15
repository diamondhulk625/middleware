package mw

import (
	"log"

	"github.com/labstack/echo/v4"
)

func StepCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		val := c.Request().Header.Get("User-Step")

		if val == "admin" {
			log.Println("User is admin")
		}

		err := next(c)
		if err != nil {
			return err
		}
		return nil
	}
}

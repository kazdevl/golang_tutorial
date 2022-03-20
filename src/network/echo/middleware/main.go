package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OneMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("pre one")
		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("after one")
		return nil
	}
}

func CancelMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("pre cancel")
		if len(c.Request().Header.Get("x-hogehoge")) != 0 {
			log.Println(c.Request().Header.Get("x-hogehoge"))
			log.Println("canceled")
			return errors.New("errored")
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("after cancel")
		return nil
	}
}

func TwoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("pre two")
		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("after two")
		return nil
	}
}

func main() {
	e := echo.New()
	e.Use(OneMiddleware, CancelMiddleware, TwoMiddleware)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success")
	})
	e.Logger.Fatal(e.Start(":1213"))
}

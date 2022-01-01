package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", func(c echo.Context) error {
		return errors.New("sample error")
	})
	log.Fatal(e.Start(":8088"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	if err := c.JSON(code, echo.NewHTTPError(code, "error that is handled")); err != nil {
		c.Logger().Error(err)
	}
}

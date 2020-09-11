package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerError(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "Server Error")
}

func ForbiddenError(c echo.Context) error {
	return c.String(http.StatusForbidden, "Forbidden Action")
}

func WrongParameters(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Parameters Error")
}

func WrongBody(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Body Error")
}
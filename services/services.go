package services

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoubleResult struct {
	Result int `json:"result"`
}

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Double(c echo.Context) error {
	number, err := strconv.Atoi(c.QueryParam("number"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "param number was invalid")
	}

	return c.JSON(http.StatusOK, &DoubleResult{Result: number * 2})
}

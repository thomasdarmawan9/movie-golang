package helpers

import (
	"h8-movies/pkg/errs"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetParamId(c echo.Context, key string) (int, errs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter id")
	}

	return id, nil
}

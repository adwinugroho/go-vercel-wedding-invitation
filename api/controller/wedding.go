package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListWishes(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"data":    "OK",
		"message": "Success!",
	})
}

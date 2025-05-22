package controller

import (
	"github.com/labstack/echo/v4"
)

func (h *WeddingController) WeddingRoutes(e *echo.Echo) {
	var wishes = e.Group("/api/wishes")
	wishes.GET("/", h.GetListWishes)
	// wishes.POST("/new", h.NewWishes)
}

package routes

import (
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/controller"
	"github.com/labstack/echo/v4"
)

func WeddingRoutes(e *echo.Echo) {
	var wishes = e.Group("/wishes")
	wishes.GET("/", controller.GetListWishes)
}

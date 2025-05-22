package handler

import (
	"net/http"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler initializes the Echo instance and handles HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world! Status OK")
	})

	routes.WeddingRoutes(e)
	// Create a new echo context and handle the request
	e.ServeHTTP(w, r)
}

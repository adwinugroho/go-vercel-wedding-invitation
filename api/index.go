package handler

import (
	"log"
	"net/http"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/controller"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/repository"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler initializes the Echo instance and handles HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world! Status OK")
	})

	// TODO: connection to DB
	repoWishes := repository.NewWishesRepository(nil)
	serviceWishes := service.NewServiceWishes(repoWishes)
	weddingController := controller.NewController(serviceWishes)

	weddingController.WeddingRoutes(e)
	e.ServeHTTP(w, r)
}

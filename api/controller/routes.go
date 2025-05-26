package controller

import (
	"net/http"
	"time"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// APIKeyMiddleware checks for valid x-api-key in request header
func (h *WeddingController) APIKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("x-api-key-wedding")
		if apiKey == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "API key header is required",
			})
		}

		// Add a small delay to prevent timing attacks
		time.Sleep(100 * time.Millisecond)
		if apiKey != config.API_KEY {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid api key",
			})
		}

		return next(c)
	}
}

func (h *WeddingController) WeddingRoutes(e *echo.Echo) {
	var wishes = e.Group("/api/wishes")

	wishes.Use(middleware.Secure())                                              // Add security headers
	wishes.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10))) // 10 req/seconds
	wishes.Use(h.APIKeyMiddleware)

	wishes.GET("/", h.GetListWishes)
	wishes.POST("/new", h.NewWishes)

	var rsvp = e.Group("/api/rsvp")

	rsvp.Use(middleware.Secure())                                              // Add security headers
	rsvp.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10))) // 10 req/seconds
	rsvp.Use(h.APIKeyMiddleware)

	rsvp.POST("/new", h.NewReservation)
}

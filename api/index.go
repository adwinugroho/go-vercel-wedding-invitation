package handler

import (
	"log"
	"net/http"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/config"
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
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_custom}","id":"${id}","remote_ip":"${remote_ip}",` + "\n" +
			`"host":"${host}","method":"${method}","uri":"${uri}","status":${status},` + "\n" +
			`"error":"${error}",` + "\n" +
			`"latency":${latency},"latency_human":"${latency_human}","bytes_in":${bytes_in},"bytes_out":${bytes_out},` + "\n" +
			`"user_agent":"${user_agent}}"` + "\n\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK! We are ready!")
	})

	// ctx := context.Background()

	// pgClient, err := config.InitPostgresConnection(ctx, config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	// if err != nil {
	// 	log.Fatal("Error while connect to client postgres:", err)
	// }

	supaClient := config.InitSupabaseConnection(config.SUPABASE_URL, config.SUPABASE_API_KEY, config.SUPABASE_PASSWORD)

	repoWishes := repository.NewWishesRepository(nil, supaClient)
	serviceWishes := service.NewServiceWishes(repoWishes)
	weddingController := controller.NewController(serviceWishes)

	weddingController.WeddingRoutes(e)
	e.ServeHTTP(w, r)
}

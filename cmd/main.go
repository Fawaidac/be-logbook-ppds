package main

import (
	"log"

	"be-logbook-ppds/app/auth"
	"be-logbook-ppds/app/jadwal"
	"be-logbook-ppds/app/user"
	"be-logbook-ppds/configs"
	"be-logbook-ppds/middleware"
	"be-logbook-ppds/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Config
	cfg := configs.LoadConfig()

	// 2. Connect to PostgreSQL
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v. Running in fallback mode.", err)
	}

	// 3. Initialize Repository, Service, and Handler
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	authService := auth.NewService(userRepo, cfg)
	authHandler := auth.NewHandler(authService)

	jadwalRepo := jadwal.NewRepository(db)
	jadwalService := jadwal.NewService(jadwalRepo)
	jadwalHandler := jadwal.NewHandler(jadwalService)

	// 4. Setup Router
	r := gin.Default()

	// CORS / Preflight middleware if needed
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(24)
			return
		}
		c.Next()
	})

	api := r.Group("/api/v1")
	{
		// Auth Endpoints
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/logout", authHandler.Logout)

			protected := authGroup.Group("")
			protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))
			{
				protected.GET("/me", authHandler.Me)
			}
		}

		// User Management CRUD (Khusus role superadmin)
		userGroup := api.Group("/users")
		userGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret), middleware.RoleMiddleware("superadmin"))
		{
			userGroup.POST("", userHandler.Create)
			userGroup.GET("", userHandler.FindAll)
			userGroup.GET("/:id", userHandler.FindByID)
			userGroup.PUT("/:id", userHandler.Update)
			userGroup.DELETE("/:id", userHandler.Delete)
		}

		// Jadwal Management Endpoints
		jadwalGroup := api.Group("/jadwals")
		{
			// Public / Protected List Events for FullCalendar
			jadwalGroup.GET("", jadwalHandler.GetEvents)

			// Edit / Create / Delete Operations (Restricted to supervisor & superadmin)
			protectedJadwal := jadwalGroup.Group("")
			protectedJadwal.Use(middleware.JWTMiddleware(cfg.JWTSecret), middleware.RoleMiddleware("supervisor", "superadmin"))
			{
				protectedJadwal.POST("", jadwalHandler.Create)
				protectedJadwal.PUT("/:id", jadwalHandler.Update)
				protectedJadwal.PATCH("/:id/dates", jadwalHandler.UpdateDates)
				protectedJadwal.DELETE("/:id", jadwalHandler.Delete)
			}
		}
	}

	// 5. Documentation & Testing UI (Tanpa mengotori kode Go)
	r.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
	r.StaticFile("/docs", "./docs/index.html")
	r.StaticFile("/swagger", "./docs/swagger.html")


	log.Printf("Server running on port :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
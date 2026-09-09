package main

import (
	"log"

	"be-logbook-ppds/app/approval"
	"be-logbook-ppds/app/auth"
	"be-logbook-ppds/app/dashboard"
	"be-logbook-ppds/app/jadwal"
	"be-logbook-ppds/app/kegiatan_ilmiah"
	"be-logbook-ppds/app/pendidikan"
	"be-logbook-ppds/app/tindakan"
	"be-logbook-ppds/app/user"
	"be-logbook-ppds/configs"
	"be-logbook-ppds/middleware"
	"be-logbook-ppds/pkg/database"
	"be-logbook-ppds/pkg/email"

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

	// 3. Initialize Mailer, Repository, Service, and Handler
	mailer := email.NewMailer(cfg)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, mailer)
	userHandler := user.NewHandler(userService)

	authService := auth.NewService(userRepo, cfg)
	authHandler := auth.NewHandler(authService)

	jadwalRepo := jadwal.NewRepository(db)
	jadwalService := jadwal.NewService(jadwalRepo)
	jadwalHandler := jadwal.NewHandler(jadwalService)

	tindakanRepo := tindakan.NewRepository(db)
	tindakanService := tindakan.NewService(tindakanRepo)
	tindakanHandler := tindakan.NewHandler(tindakanService)

	kegiatanRepo := kegiatan_ilmiah.NewRepository(db)
	bimbinganRepo := kegiatan_ilmiah.NewBimbinganRepository(db)
	kegiatanService := kegiatan_ilmiah.NewService(kegiatanRepo, bimbinganRepo)
	kegiatanHandler := kegiatan_ilmiah.NewHandler(kegiatanService)

	kompetensiRepo := pendidikan.NewKompetensiRepository(db)
	kompetensiService := pendidikan.NewKompetensiService(kompetensiRepo)

	rotasiRepo := pendidikan.NewRotasiRepository(db)
	rotasiService := pendidikan.NewRotasiService(rotasiRepo)

	miniCexRepo := pendidikan.NewMiniCexRepository(db)
	miniCexService := pendidikan.NewMiniCexService(miniCexRepo)

	dopsRepo := pendidikan.NewDopsRepository(db)
	dopsService := pendidikan.NewDopsService(dopsRepo)

	seminarRepo := pendidikan.NewSeminarRepository(db)
	seminarService := pendidikan.NewSeminarService(seminarRepo)

	cbdRepo := pendidikan.NewCbdRepository(db)
	cbdService := pendidikan.NewCbdService(cbdRepo)

	pendidikanHandler := pendidikan.NewHandler(kompetensiService, rotasiService, miniCexService, dopsService, seminarService, cbdService)

	// Approval Service - wraps existing repos
	approvalTindakanRepo := &approval.TindakanRepoAdapter{DB: db}
	approvalKegiatanRepo := &approval.KegiatanIlmiahRepoAdapter{DB: db}
	approvalPendidikanRepo := &approval.PendidikanEvaluasiRepoAdapter{DB: db}

	approvalService := approval.NewService(approvalTindakanRepo, approvalKegiatanRepo, approvalPendidikanRepo)
	approvalHandler := approval.NewHandler(approvalService)

	dashboardRepo := dashboard.NewRepository(db)
	dashboardService := dashboard.NewService(dashboardRepo, kompetensiRepo)
	dashboardHandler := dashboard.NewHandler(dashboardService)

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
			authGroup.POST("/register", userHandler.Register)
			authGroup.POST("/pre-register", userHandler.PreRegisterCheck)

			protected := authGroup.Group("")
			protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))
			{
				protected.POST("/logout", authHandler.Logout)
				protected.GET("/me", authHandler.Me)
			}
		}

		// User Management CRUD & Verifikasi (Khusus role admin)
		userGroup := api.Group("/users")
		userGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret), middleware.RoleMiddleware("admin"))
		{
			userGroup.POST("", userHandler.Create)
			userGroup.GET("", userHandler.FindAll)
			userGroup.GET("/registrations", userHandler.GetRegistrations)
			userGroup.POST("/registrations/:id/approve", userHandler.ApproveRegistration)
			userGroup.POST("/registrations/:id/reject", userHandler.RejectRegistration)
			userGroup.GET("/:id", userHandler.FindByID)
			userGroup.PUT("/:id", userHandler.Update)
			userGroup.DELETE("/:id", userHandler.Delete)
		}

		// Profile / Master Data endpoints (user yang sedang login, khusus residen dari FE)
		profileGroup := api.Group("/users/profile")
		profileGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			profileGroup.GET("", userHandler.GetProfile)
			profileGroup.PUT("", userHandler.UpdateProfile)
		}

		workHistoryGroup := api.Group("/users/work-histories")
		workHistoryGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			workHistoryGroup.GET("", userHandler.GetWorkHistories)
			workHistoryGroup.POST("", userHandler.CreateWorkHistory)
			workHistoryGroup.PUT("/:id", userHandler.UpdateWorkHistory)
			workHistoryGroup.DELETE("/:id", userHandler.DeleteWorkHistory)
		}

		// Jadwal Management Endpoints
		jadwalGroup := api.Group("/jadwals")
		jadwalGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			jadwalGroup.GET("", jadwalHandler.GetEvents)

			protectedJadwal := jadwalGroup.Group("")
			protectedJadwal.Use(middleware.RoleMiddleware("supervisor", "admin", "residen"))
			{
				protectedJadwal.POST("", jadwalHandler.Create)
				protectedJadwal.PUT("/:id", jadwalHandler.Update)
				protectedJadwal.PATCH("/:id/dates", jadwalHandler.UpdateDates)
				protectedJadwal.DELETE("/:id", jadwalHandler.Delete)
			}
		}

		// Tindakan Logbook Endpoints
		tindakanGroup := api.Group("/tindakans")
		tindakanGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			tindakanGroup.GET("", tindakanHandler.GetSummary)
			tindakanGroup.GET("/getdpjp", tindakanHandler.GetDPJP)
			tindakanGroup.GET("/by-department", tindakanHandler.GetByDepartment)
			tindakanGroup.GET("/by-supervisor/:name", tindakanHandler.GetBySupervisor)
			tindakanGroup.GET("/:id", tindakanHandler.GetByID)
			tindakanGroup.POST("", tindakanHandler.Create)
			tindakanGroup.PUT("/:id", tindakanHandler.Update)
			tindakanGroup.POST("/:id/send", tindakanHandler.Send)
			tindakanGroup.DELETE("/:id", tindakanHandler.Delete)
		}

		// Kegiatan Ilmiah & Bimbingan Penelitian Endpoints
		kegiatanGroup := api.Group("/kegiatan-ilmiah")
		kegiatanGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			kegiatanGroup.GET("", kegiatanHandler.GetIndex)
			kegiatanGroup.POST("", kegiatanHandler.Create)
			kegiatanGroup.DELETE("/:id", kegiatanHandler.Delete)
		}

		bimbinganGroup := api.Group("/bimbingans")
		bimbinganGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			bimbinganGroup.GET("", kegiatanHandler.GetBimbinganIndex)
			bimbinganGroup.POST("", kegiatanHandler.CreateBimbingan)
		}

		// Pendidikan Endpoints
		kompetensiGroup := api.Group("/pendidikan/kompetensi")
		kompetensiGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			kompetensiGroup.GET("", pendidikanHandler.GetKompetensi)
			kompetensiGroup.POST("", pendidikanHandler.CreateKompetensi)
			kompetensiGroup.DELETE("/:id", pendidikanHandler.DeleteKompetensi)
		}

		rotasiGroup := api.Group("/pendidikan/rotasi")
		rotasiGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			rotasiGroup.GET("", pendidikanHandler.GetRotasi)
			rotasiGroup.POST("", pendidikanHandler.CreateRotasi)
			rotasiGroup.DELETE("/:id", pendidikanHandler.DeleteRotasi)
		}

		miniCexGroup := api.Group("/pendidikan/mini-cex")
		miniCexGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			miniCexGroup.GET("", pendidikanHandler.GetMiniCex)
			miniCexGroup.POST("", pendidikanHandler.CreateMiniCex)
			miniCexGroup.DELETE("/:id", pendidikanHandler.DeleteMiniCex)
		}

		dopsGroup := api.Group("/pendidikan/dops")
		dopsGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			dopsGroup.GET("", pendidikanHandler.GetDops)
			dopsGroup.POST("", pendidikanHandler.CreateDops)
			dopsGroup.DELETE("/:id", pendidikanHandler.DeleteDops)
		}

		seminarGroup := api.Group("/pendidikan/seminar")
		seminarGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			seminarGroup.GET("", pendidikanHandler.GetSeminar)
			seminarGroup.POST("", pendidikanHandler.CreateSeminar)
			seminarGroup.DELETE("/:id", pendidikanHandler.DeleteSeminar)
		}

		cbdGroup := api.Group("/pendidikan/cbd")
		cbdGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			cbdGroup.GET("", pendidikanHandler.GetCbd)
			cbdGroup.POST("", pendidikanHandler.CreateCbd)
			cbdGroup.DELETE("/:id", pendidikanHandler.DeleteCbd)
		}

		// Approval Endpoints
		approvalGroup := api.Group("/approval")
		approvalGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret), middleware.RoleMiddleware("supervisor", "admin"))
		{
			approvalGroup.GET("/menunggu", approvalHandler.GetMenunggu)
			approvalGroup.GET("/disetujui", approvalHandler.GetDisetujui)
			approvalGroup.GET("/ditolak", approvalHandler.GetDitolak)

			// Tindakan Approval
			approvalGroup.POST("/tindakans/:id/approve", approvalHandler.ApproveTindakan)
			approvalGroup.POST("/tindakans/:id/reject", approvalHandler.RejectTindakan)

			// Kegiatan Ilmiah Approval
			approvalGroup.POST("/kegiatan-ilmiah/:id/approve", approvalHandler.ApproveKegiatanIlmiah)
			approvalGroup.POST("/kegiatan-ilmiah/:id/reject", approvalHandler.RejectKegiatanIlmiah)

			// Aktivitas Klinik Approval

			// Pendidikan Evaluasi Approval
			approvalGroup.POST("/pendidikan-evaluasi/:id/approve", approvalHandler.ApprovePendidikanEvaluasi)
			approvalGroup.POST("/pendidikan-evaluasi/:id/reject", approvalHandler.RejectPendidikanEvaluasi)
		}

		// Dashboard Endpoints
		dashboardGroup := api.Group("/dashboard")
		dashboardGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
		{
			dashboardGroup.GET("", dashboardHandler.GetDashboardSummary)
			dashboardGroup.GET("/residen", dashboardHandler.GetResidentDashboard)
			dashboardGroup.GET("/admin", middleware.RoleMiddleware("admin"), dashboardHandler.GetAdminDashboard)
			dashboardGroup.GET("/supervisor", middleware.RoleMiddleware("supervisor"), dashboardHandler.GetSupervisorDashboard)
			dashboardGroup.GET("/laporan", dashboardHandler.GetLaporanSummary)
		}
	}

	// 5. Documentation & Uploads UI
	r.Static("/uploads", "./uploads")
	r.Static("/api/uploads", "./uploads")
	r.Static("/api/v1/uploads", "./uploads")
	r.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
	r.StaticFile("/docs", "./docs/index.html")
	r.StaticFile("/swagger", "./docs/swagger.html")

	log.Printf("Server running on port :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

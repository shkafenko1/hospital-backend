package app

import (
	"log"

	"hospital-backend/internal/configs"
	"hospital-backend/internal/handlers"
	"hospital-backend/internal/middleware"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/service"
	"hospital-backend/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func App() {
	logger.InitLogger()

	db := configs.GetDB()
	defer db.Close()

	if err := configs.RunMigrations(db, "migrations"); err != nil {
		log.Fatalf("migrations failed: %v", err)
	}
	if err := configs.SeedDefaultUsers(db); err != nil {
		log.Fatalf("seeding failed: %v", err)
	}

	app := fiber.New()

	// API documentation
	app.Get("/openapi.yaml", handlers.OpenAPISpec)
	app.Get("/docs", handlers.Redoc)

	api := app.Group("/api")

	// Initialize repositories
	patientRepo := repository.NewPatientRepository(db)
	doctorRepo := repository.NewDoctorRepository(db)
	wardRepo := repository.NewWardRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)
	hospitalizationRepo := repository.NewHospitalizationRepository(db)
	visitRepo := repository.NewVisitRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	patientService := service.NewPatientService(patientRepo)
	doctorService := service.NewDoctorService(doctorRepo)
	wardService := service.NewWardService(wardRepo)
	medicineService := service.NewMedicineService(medicineRepo)
	hospitalizationService := service.NewHospitalizationService(hospitalizationRepo)
	visitService := service.NewVisitService(visitRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	patientHandler := handlers.NewPatientHandler(patientService)
	doctorHandler := handlers.NewDoctorHandler(doctorService)
	wardHandler := handlers.NewWardHandler(wardService)
	medicineHandler := handlers.NewMedicineHandler(medicineService)
	hospitalizationHandler := handlers.NewHospitalizationHandler(hospitalizationService)
	visitHandler := handlers.NewVisitHandler(visitService)
	authHandler := handlers.NewAuthHandler(userService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(userService)

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// Patient routes (protected)
	patients := api.Group("/patients")
	patients.Use(authMiddleware.Authenticate)
	patients.Post("/", patientHandler.Create)
	patients.Get("/", patientHandler.GetAll)
	patients.Get("/:id", patientHandler.GetByID)
	patients.Put("/:id", patientHandler.Update)
	patients.Delete("/:id", patientHandler.Delete)

	// Doctor routes (protected - admin only for write operations)
	doctors := api.Group("/doctors")
	doctors.Use(authMiddleware.Authenticate)
	doctors.Get("/", doctorHandler.GetAll)
	doctors.Get("/:id", doctorHandler.GetByID)
	doctors.Post("/", authMiddleware.RequireAdmin(), doctorHandler.Create)
	doctors.Put("/:id", authMiddleware.RequireAdmin(), doctorHandler.Update)
	doctors.Delete("/:id", authMiddleware.RequireAdmin(), doctorHandler.Delete)

	// Ward routes (protected - admin only for write operations)
	wards := api.Group("/wards")
	wards.Use(authMiddleware.Authenticate)
	wards.Get("/", wardHandler.GetAll)
	wards.Get("/:id", wardHandler.GetByID)
	wards.Post("/", authMiddleware.RequireAdmin(), wardHandler.Create)
	wards.Put("/:id", authMiddleware.RequireAdmin(), wardHandler.Update)
	wards.Delete("/:id", authMiddleware.RequireAdmin(), wardHandler.Delete)

	// Medicine routes (protected - admin only for write operations)
	medicines := api.Group("/medicines")
	medicines.Use(authMiddleware.Authenticate)
	medicines.Get("/", medicineHandler.GetAll)
	medicines.Get("/:id", medicineHandler.GetByID)
	medicines.Post("/", authMiddleware.RequireAdmin(), medicineHandler.Create)
	medicines.Put("/:id", authMiddleware.RequireAdmin(), medicineHandler.Update)
	medicines.Delete("/:id", authMiddleware.RequireAdmin(), medicineHandler.Delete)

	// Hospitalization routes (protected)
	hospitalizations := api.Group("/hospitalizations")
	hospitalizations.Use(authMiddleware.Authenticate)
	hospitalizations.Post("/", hospitalizationHandler.Create)
	hospitalizations.Get("/", hospitalizationHandler.GetAll)
	hospitalizations.Get("/:id", hospitalizationHandler.GetByID)
	hospitalizations.Put("/:id", hospitalizationHandler.Update)
	hospitalizations.Delete("/:id", hospitalizationHandler.Delete)

	// Visit routes (protected)
	visits := api.Group("/visits")
	visits.Use(authMiddleware.Authenticate)
	visits.Post("/", visitHandler.Create)
	visits.Get("/", visitHandler.GetAll)
	visits.Get("/:id", visitHandler.GetByID)
	visits.Put("/:id", visitHandler.Update)
	visits.Delete("/:id", visitHandler.Delete)

	app.Listen(":8000")
}

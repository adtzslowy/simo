package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adtzslowy/simo/internal/auth"
	"github.com/adtzslowy/simo/internal/config"
	"github.com/adtzslowy/simo/internal/database"
	"github.com/adtzslowy/simo/internal/handler"
	"github.com/adtzslowy/simo/internal/middleware"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/adtzslowy/simo/internal/router"
	"github.com/adtzslowy/simo/internal/service"
)

func main() {
	// =========================================================
	// CONFIG
	// =========================================================

	cfg := config.Load()

	// =========================================================
	// DATABASE
	// =========================================================

	db, err := database.NewPostgresPool(database.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// =========================================================
	// REPOSITORY
	// =========================================================

	organizationTypeRepository := repository.NewOrganizationTypeRepository(db)

	organizationRepository := repository.NewOrganizationRepository(db)

	organizationPeriodRepository := repository.NewOrganizationPeriodRepository(db)

	userRepository := repository.NewUserRepository(db)

	memberPositionRepository := repository.NewMemberPositionRepository(db)

	// =========================================================
	// AUTH
	// =========================================================

	jwtService := auth.NewJWTService(
		os.Getenv("JWT_SECRET"),
		24*time.Hour,
	)

	authMiddleware := middleware.NewAuthMiddleware(
		jwtService,
	)

	// =========================================================
	// SERVICE
	// =========================================================

	organizationTypeService := service.NewOrganizationTypeService(
		organizationTypeRepository,
	)

	organizationService := service.NewOrganizationService(
		organizationRepository,
	)

	organizationPeriodService := service.NewOrganizationPeriodService(
		organizationPeriodRepository,
	)

	authService := service.NewAuthService(
		userRepository,
		jwtService,
	)

	memberPositionService := service.NewMemberPositionService(
		memberPositionRepository,
	)

	// =========================================================
	// HANDLER
	// =========================================================

	organizationTypeHandler := handler.NewOrganizationTypeHandler(
		organizationTypeService,
	)

	organizationHandler := handler.NewOrganizationHandler(
		organizationService,
	)

	organizationPeriodHandler := handler.NewOrganizationPeriodHandler(
		organizationPeriodService,
	)

	authHandler := handler.NewAuthHandler(
		authService,
	)

	memberPositionHandler := handler.NewMemberPositionHandler(
		memberPositionService,
	)

	// =========================================================
	// ROUTER
	// =========================================================

	r := router.New(
		organizationTypeHandler,
		organizationHandler,
		organizationPeriodHandler,
		memberPositionHandler,
		authHandler,
		authMiddleware,
	)

	// =========================================================
	// SERVER
	// =========================================================

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	log.Println("server running on :" + cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

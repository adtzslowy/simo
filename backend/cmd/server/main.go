package main

import (
	"log"
	"net/http"

	"github.com/adtzslowy/simo/internal/config"
	"github.com/adtzslowy/simo/internal/database"
	"github.com/adtzslowy/simo/internal/handler"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/adtzslowy/simo/internal/router"
	"github.com/adtzslowy/simo/internal/service"
)

func main() {
	cfg := config.Load()

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

	// Repository
	organizationTypeRepository := repository.NewOrganizationTypeRepository(db)
	organizationRepository := repository.NewOrganizationRepository(db)
	organizationPeriodRepository := repository.NewOrganizationPeriodRepository(db)

	// Service
	organizationTypeService := service.NewOrganizationTypeService(
		organizationTypeRepository,
	)
	organizationService := service.NewOrganizationService(
		organizationRepository,
	)
	organizationPeriodService := service.NewOrganizationPeriodService(organizationPeriodRepository)

	// Handler
	organizationTypeHandler := handler.NewOrganizationTypeHandler(
		organizationTypeService,
	)
	organizationHandler := handler.NewOrganizationHandler(
		organizationService,
	)
	organizationPeriodHandler := handler.NewOrganizationPeriodHandler(organizationPeriodService)

	// Router
	r := router.New(
		organizationTypeHandler,
		organizationHandler,
		organizationPeriodHandler,
	)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	log.Println("server running on :" + cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

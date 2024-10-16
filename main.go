package main

import (
	"log"

	"github.com/Rezist5/Ai-BeeReg/config"
	"github.com/Rezist5/Ai-BeeReg/migrations"
	"github.com/Rezist5/Ai-BeeReg/routes"
	"github.com/Rezist5/Ai-BeeReg/seeders"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	migrations.RunMigrations(db)
	if err := seeders.Run(db); err != nil {
		log.Fatalf("Failed to run seeders: %v", err)
	}
	router := gin.Default()

	routes.SetupRoutes(router, db)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

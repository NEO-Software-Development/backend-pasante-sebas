package main

import (
	"log"

	"github.com/JuanPO17/myfirstgo/internal/config"
	"github.com/JuanPO17/myfirstgo/internal/database"
	"github.com/JuanPO17/myfirstgo/internal/handler"
	"github.com/JuanPO17/myfirstgo/internal/repository"
	"github.com/JuanPO17/myfirstgo/internal/router"
	_ "github.com/JuanPO17/myfirstgo/docs"
)

// @title Go Task Manager API
// @version 1.0
// @description This is a sample server for a task manager application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to the database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the database
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository
	taskRepo := repository.NewTaskRepository(db)

	// Initialize handler
	taskHandler := handler.NewTaskHandler(taskRepo)

	// Setup router
	r := router.SetupRouter(taskHandler)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

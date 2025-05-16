package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/broot5/formula-place/server/internal/config"
	"github.com/broot5/formula-place/server/internal/database"
	"github.com/broot5/formula-place/server/internal/handlers"
	"github.com/broot5/formula-place/server/internal/repositories"
	"github.com/broot5/formula-place/server/internal/routes"
	"github.com/broot5/formula-place/server/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbPool, err := database.New(context.Background(), cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	formulaRepo := repositories.NewFormulaRepository(dbPool)
	formulaService := services.NewFormulaService(formulaRepo)
	formulaHandler := handlers.NewFormulaHandler(formulaService)

	deps := &routes.RouterDeps{
		FormulaHandler: formulaHandler,
	}

	r := routes.NewRouter(deps)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.ServerPort),
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on port %d", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

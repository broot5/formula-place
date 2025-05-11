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
	"github.com/broot5/formula-place/server/internal/services"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		config := huma.DefaultConfig("Formula Place API", "1.0.0")
		config.Servers = []*huma.Server{
			{URL: "http://localhost:3000/api"},
		}

		api := humachi.New(r, config)

		huma.Post(api, "/formulas", formulaHandler.CreateFormula)
		huma.Get(api, "/formulas/{id}", formulaHandler.GetFormula)
		huma.Patch(api, "/formulas/{id}", formulaHandler.UpdateFormula)
		huma.Delete(api, "/formulas/{id}", formulaHandler.DeleteFormula)
		huma.Get(api, "/formulas", formulaHandler.GetAllFormulas)
	})

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

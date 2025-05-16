package routes

import (
	"time"

	"github.com/broot5/formula-place/server/internal/handlers"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouterDeps struct {
	FormulaHandler *handlers.FormulaHandler
}

func NewRouter(deps *RouterDeps) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		config := huma.DefaultConfig("Formula Place API", "1.0.0")
		config.Servers = []*huma.Server{
			{URL: "http://localhost:3000/api"},
		}

		api := humachi.New(r, config)

		huma.Post(api, "/formulas", deps.FormulaHandler.CreateFormula)
		huma.Get(api, "/formulas/{id}", deps.FormulaHandler.GetFormula)
		huma.Patch(api, "/formulas/{id}", deps.FormulaHandler.UpdateFormula)
		huma.Delete(api, "/formulas/{id}", deps.FormulaHandler.DeleteFormula)
		huma.Get(api, "/formulas", deps.FormulaHandler.GetAllFormulas)
	})

	return router
}

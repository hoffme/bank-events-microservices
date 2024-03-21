package api

import (
	"net/http"
	"time"

	"github.com/hoffme/backend-transactions/internal/domain"

	"github.com/hoffme/backend-transactions/internal/application/api/controllers"
	"github.com/hoffme/backend-transactions/internal/application/api/generated"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target generated --clean -package generated doc.yml

func Handler(deps domain.Dependencies) (http.Handler, error) {
	ctr := controllers.Controllers{Deps: deps}
	handler, err := generated.NewServer(ctr, generated.WithPathPrefix("/api/v1"))
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Compress(5))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	r.Mount("/api/v1", handler)

	return r, nil
}

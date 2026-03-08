package router

import (
	_ "kt_project/docs"
	"kt_project/internal/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	// creating router
	r := chi.NewRouter()

	// adding middlewares
	r.Use(middleware.Logger)    // writes every request to console
	r.Use(middleware.Recoverer) // keeps server up

	r.Get("/", h.Main)
	// for docs
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// when agent posts, dumpmetric is called
	r.Post("/update", h.DumpStat)
	// test router to check in browser
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r
}

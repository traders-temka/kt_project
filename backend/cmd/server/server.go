package server

import (
	"fmt"
	"backend/internal/repository"
	"backend/internal/server/handlers"
	"backend/internal/server/router"
	"log"
	"net/http"
)

func Run() error {
	// initializing storage to mem (later to SQLite)
	repo, err := repository.NewSqlStorage("metrics.db")
	if err != nil {
		return fmt.Errorf("failed to create reposql: %w", err)
	}

	// init handlers and give them repo
	h := &handlers.Handler{Repo: repo}

	// creating router
	r := router.NewRouter(h)

	log.Println("server is up on :8080")
	// up port for listening
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error starting server: %s", err)
	}

	return nil
}

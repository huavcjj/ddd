package main

import (
	"context"
	"database/sql"
	"ddd/repository/sample/repository"
	"ddd/repository/sample/service"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./sample.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT UNIQUE NOT NULL
	);`
	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

}
func main() {

	server := &http.Server{
		Addr:    ":8081",
		Handler: setupRoutes(),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("API server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	select {
	case err := <-serverErrors:
		log.Fatalf("Server error: %v", err)

	case sig := <-shutdown:
		log.Printf("Start shutdown... Signal: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)

		if err := server.Close(); err != nil {
			log.Printf("Forced shutdown failed: %v", err)
		}
	}

	log.Printf("Server shutdown complete")
}

func setupRoutes() http.Handler {
	userRepository := repository.NewUserRepository(db)
	svc := service.NewUserService(userRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		name := body["name"]
		if name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		user, err := svc.CreateUser(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{
			"user_id": user.ID,
			"name":    user.Name,
		})
	})
	return mux
}

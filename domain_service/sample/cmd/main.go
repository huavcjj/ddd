package main

import (
	"context"
	"ddd/domain_service/sample/service"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	server := &http.Server{
		Addr:    ":8080",
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
	mux := http.NewServeMux()

	svc := service.NewUserService()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user, err := svc.CreateUser("test_user_name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"user_id": user.UserID.String(),
			"name":    user.UserName.String(),
		})
	})
	return mux
}

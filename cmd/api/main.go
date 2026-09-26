package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmed-wassim/wassimo-catalog/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("error happened %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("catalog listening on :%s", cfg.PORT)
	if err := http.ListenAndServe(":"+cfg.PORT, mux); err != nil {
		log.Fatal(err)
	}
}

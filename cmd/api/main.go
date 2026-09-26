package main

import (
	"log"
	"net/http"

	"github.com/ahmed-wassim/wassimo-catalog/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	log.Printf("catalog listening on :%s", cfg.PORT)
	if err := r.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}
}

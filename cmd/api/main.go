package main

import (
	"log"
	"net/http"

	"github.com/ahmed-wassim/wassimo-catalog/internal/config"
	"github.com/ahmed-wassim/wassimo-catalog/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	r := gin.Default()

	r.GET("/ready", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"ready": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ready": true})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	h := handlers.Handler{DB: db}
	r.GET("/restaurants", h.List)
	r.GET("/restaurants/:id/branches", h.ListBranches)
	r.POST("/restaurants/:id/branches", h.CreateBranch)

	log.Printf("catalog listening on :%s", cfg.PORT)
	if err := r.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}
}

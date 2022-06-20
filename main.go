package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-prices-tracker-rest-api/controllers"
	"net/http"
	"os"
	"time"
)

var (
	ApiVersion  = "/v1"
	defaultPort = "8080"
)

func main() {

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.GET(ApiVersion+"/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"version": ApiVersion,
		})
	})

	engine.GET(ApiVersion+"/products", controllers.ListProducts)

	if envPort := os.Getenv("PORT"); envPort != "" {
		defaultPort = envPort
	}

	engine.Run(":" + defaultPort)
}

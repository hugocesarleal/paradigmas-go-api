package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"trab-final/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/sequential", handlers.HandleSequential)
		api.GET("/parallel", handlers.HandleParallel)
		api.GET("/benchmark", handlers.HandleBenchmark)
		api.GET("/graph", handlers.HandleGraph)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

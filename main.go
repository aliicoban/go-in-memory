package main

import (
	routes "github.com/alicobanserver/routes"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"log"
	"os"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(router.Run(":4444"))
	} else {
		log.Fatal(router.Run(":" + port))
	}
}

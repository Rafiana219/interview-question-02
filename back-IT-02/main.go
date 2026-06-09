package main

import (
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/config"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Rafiana219/interview-question-02/back-IT-02/docs"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

// @title IT-02 Authentication API
// @version 1.0
// @description Authentication System
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	err := godotenv.Load(".env")
	// utils.InitJWT()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r)

	// r.GET("/test", handlers.TestAPI)
	r.Run(":8080")
}

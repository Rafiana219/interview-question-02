package routes

import (
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/handlers"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}
	protected := api.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", handlers.Profile)
	}

}

package handlers

import (
	"net/http"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/repositories"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/services"

	"github.com/gin-gonic/gin"
)

var authService = services.NewAuthService(
	repositories.NewUserRepository(),
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register User
// @Description Create new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register Request"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := authService.Register(
		req.Username,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
	})
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	token, err :=
		authService.Login(
			req.Username,
			req.Password,
		)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Profile(c *gin.Context) {

	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user, err := authService.GetProfile(
		userId.(uint),
	)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
	})
}

func TestAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Go Backend",
	})
}

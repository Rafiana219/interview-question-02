package handlers

import (
	"net/http"

	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/config"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/models"
	"github.com/Rafiana219/interview-question-02/back-IT-02/internal/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
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
			"message": "invalid request",
		})

		return
	}

	hash, _ := utils.HashPassword(req.Password)

	user := models.User{
		Username: req.Username,
		Password: hash,
	}

	config.DB.Create(&user)

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

	var req RegisterRequest

	c.ShouldBindJSON(&req)

	var user models.User

	result := config.DB.
		Where("username = ?", req.Username).
		First(&user)

	if result.Error != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid username",
		})

		return
	}

	if !utils.CheckPassword(
		user.Password,
		req.Password,
	) {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid password",
		})

		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Profile(c *gin.Context) {

	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	var user models.User

	config.DB.First(&user, userId)

	c.JSON(200, gin.H{
		"username": user.Username,
	})
}

func TestAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Go Backend",
	})
}

package controllers

import (
	"net/http"

	"com.lopster-pos/models"
	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true "login to access your account"
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request body"})
		return
	}

	if req.Phone == "" ||
		req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill all required fields"})
		return
	}

	res, err := services.LoginService(req)
	if err != nil {
		println(http.StatusUnauthorized, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid phone or password"})
		return
	}

	c.JSON(http.StatusOK, res)
}

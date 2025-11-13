package controllers

import (
	"log"
	"net/http"

	"com.lopster-pos/models"
	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Register request body"
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request body"})
		return
	}

	if req.FirstName == "" || req.LastName == "" ||
		req.Gender == "" || req.Phone == "" ||
		req.Role == "" || req.PasswordHash == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please fill all required fields",
		})
		return
	}

	res, err := services.RegisterService(req)
	if err != nil {
		if err.Error() == "User with this phone number already exists, please login" {
			log.Print(http.StatusBadRequest, err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User with this phone number already exists, please login",
			})
			return
		}
		log.Print(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

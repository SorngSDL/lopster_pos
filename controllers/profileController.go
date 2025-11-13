package controllers

import (
	"log"
	"net/http"

	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Tags Users
// @Produce json
// @Router /api/users/get-profile [get]
func GetProfile(c *gin.Context) {
	uid, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
		return
	}

	userID, _ := uid.(string)

	profile, err := services.GetProfileService(userID)
	if err != nil {
		log.Print(http.StatusNotFound, err)

		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		log.Print(http.StatusInternalServerError, err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

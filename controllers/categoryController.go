package controllers

import (
	"net/http"

	"com.lopster-pos/models"
	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Tags Menu
// @Accept json
// @Produce json
// @Param data body models.CategoryRequest true "Create category request body"
// @Router /api/menu/create-category [post]
func CreateCategory(c *gin.Context) {
	var req models.CategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.CategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": " category is required"})
		return
	}

	resp, err := services.CategoryService(req)
	if err != nil {
		if err.Error() == "category already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create category"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

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
// @Param data body models.MenuTypeRequest true "Menu type request body"
// @Router /api/menu/menu-type [post]
func CreateMenuType(c *gin.Context) {
	var req models.MenuTypeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.MenuTypeName == "" || req.MenuTypeName == "string" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "menu type name is required"})
		return
	}

	resp, err := services.MenuTypeService(req)
	if err != nil {
		if err.Error() == "menu type already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu type already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create menu type"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

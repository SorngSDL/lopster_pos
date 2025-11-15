package controllers

import (
	"net/http"

	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Summary Get all menus
// @Tags Menu
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MenuResponse
// @Router /api/menu/list [get]
func ListMenu(c *gin.Context) {
	resp, err := services.ListMenuService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch menus"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// controllers/menu_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"com.lopster-pos/models"
	"com.lopster-pos/services"
	"github.com/gin-gonic/gin"
)

// @Summary Edit menu
// @Tags Menu
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path string true "Menu ID"
// @Param menuNameEn formData string true "Menu name EN"
// @Param menuNameLo formData string true "Menu name LO"
// @Param categoryName formData string true "Category name"
// @Param price formData number true "Price"
// @Param unit formData string true "Unit"
// @Param promotionName formData string false "Promotion name"
// @Param promotionPrice formData number false "Promotion price"
// @Param menuImage formData file false "Menu image (optional, send to replace)"
// @Success 200 {object} models.MenuResponse
// @Router /api/menu/edit-menu/{id} [put]
func EditMenu(c *gin.Context) {
	id := c.Param("id")

	var req models.MenuRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var imageURLPtr *string
	file, err := c.FormFile("menuImage")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save image"})
			return
		}
		tmp := "/uploads/" + filename
		imageURLPtr = &tmp
	}

	resp, err := services.UpdateMenuService(id, req, imageURLPtr)
	if err != nil {
		switch err.Error() {
		case "invalid menu id":
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid menu id"})
			return
		case "menu not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update menu"})
			return
		}
	}

	c.JSON(http.StatusOK, resp)
}

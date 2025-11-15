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

// @Summary Create menu
// @Tags Menu
// @Accept multipart/form-data
// @Produce json
// @Param menuNameEn formData string true "Menu name EN"
// @Param menuNameLo formData string true "Menu name LO"
// @Param categoryName formData string true "Category name"
// @Param price formData number true "Price"
// @Param unit formData string true "Unit"
// @Param promotionName formData string false "Promotion name"
// @Param promotionPrice formData number false "Promotion price"
// @Param menuImage formData file false "Menu image"
// @Router /api/menu/create-menu [post]
func CreateMenu(c *gin.Context) {
	var req models.MenuRequest

	// สำหรับ debug ให้โชว์ error จริง
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("menuImage")
	imageURL := ""
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save image"})
			return
		}
		imageURL = "/uploads/" + filename
	}

	resp, err := services.CreateMenuService(req, imageURL)
	if err != nil {
		if err.Error() == "menu already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create menu"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

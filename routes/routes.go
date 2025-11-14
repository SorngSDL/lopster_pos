package routes

import (
	"com.lopster-pos/controllers"
	"com.lopster-pos/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	auth := r.Group("/api")
	{
		auth.POST("/auth/register", controllers.Register)
		auth.GET("/users/get-profile", middlewares.AuthMiddleware(), controllers.GetProfile)
		auth.POST("/auth/login", controllers.Login)
		auth.POST("/menu/menu-type", middlewares.AuthMiddleware(), controllers.CreateMenuType)

	}
}

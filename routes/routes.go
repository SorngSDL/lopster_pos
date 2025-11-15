package routes

import (
	"com.lopster-pos/controllers"
	"com.lopster-pos/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	apiRoute := r.Group("/api")
	{
		apiRoute.POST("/auth/register", controllers.Register)
		apiRoute.GET("/users/get-profile", middlewares.AuthMiddleware(), controllers.GetProfile)
		apiRoute.POST("/auth/login", controllers.Login)
		apiRoute.POST("/menu/create-category", middlewares.AuthMiddleware(), controllers.CreateCategory)
		apiRoute.POST("/menu/create-menu", middlewares.AuthMiddleware(), controllers.CreateMenu)

	}
}

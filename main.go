package main

import (
	"log"

	"com.lopster-pos/db"
	"com.lopster-pos/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "com.lopster-pos/docs"
)

// @title Lopster POS API
// @version 1.0
// @description Lopster POS backend
// @host localhost:8080
// @BasePath /
func main() {
	db.ConnectDB()

	r := gin.Default()
	routes.Routes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server started on :8080")
	r.Run(":8080")
}

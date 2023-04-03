package main

import (
	"jwt-authentication-golang/internal/controllers"
	"jwt-authentication-golang/internal/database"
	"jwt-authentication-golang/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	dsn := "host=localhost user=postgres password=password dbname=jwt_test_db port=5432 sslmode=disable"
	database.Connect(dsn)
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

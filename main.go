package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan2bd/jwt-authentication-golang/controllers"
	"github.com/raihan2bd/jwt-authentication-golang/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", controllers.Validate)

	router.Run()
}

package routes

import (
	userController "go-mnc/controller/userController"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("/register", userController.RegisterHandler) 
		userGroup.GET("/login", userController.LoginHandler) 
	}
}

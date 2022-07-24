package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/controller"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	service "github.com/gondsuryaprakash/gondpariwar/service/login"
)

func UserHandler(router *gin.Engine) {
	funcName := "handler.UserHandler"
	logger.I(funcName)
	auth := router.Group("/user")
	{
		auth.POST("/login", controller.PostLogin)
		auth.POST("/register", controller.AddUser)
		auth.POST("/forgotpassword", service.IsUserExist(), controller.PostForgotPassword)
		// auth.PUT("/update", controller.UpdateUser)
		// auth.DELETE("/delete/:id", controller.DeleteUser)
	}

	userDetails := router.Group("/me", service.Authorise())
	{
		userDetails.GET("/", controller.GetUserById)
		// userDetails.GET("/:mobile", controller.GetUserByMobile)
	}
}

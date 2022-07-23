package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/controller"
	"github.com/gondsuryaprakash/gondpariwar/logger"
)

func UserHandler(router *gin.Engine) {
	funcName := "handler.UserHandler"
	logger.I(funcName)
	auth := router.Group("/auth", controller.PostLogin)
	{
		auth.POST("/login")
		auth.POST("/register", controller.AddUser)
		// auth.PUT("/update", controller.UpdateUser)
		// auth.DELETE("/delete/:id", controller.DeleteUser)
	}

	userDetails := router.Group("/userinfo")
	{
		userDetails.GET("/", controller.GetUserById)
		// userDetails.GET("/:mobile", controller.GetUserByMobile)
	}
}

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gondsuryaprakash/gondpariwar/database"
	"github.com/gondsuryaprakash/gondpariwar/handler"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/utils"
)

func init() {
	funcName := "main.init"
	database.Connection()
	logger.I(funcName)
}

func main() {
	funcName := "main.main"
	logger.I(funcName)
	ConstPORT := utils.GetConfigValue("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	handler.UserHandler(router)
	// User Handler
	router.Run(":" + ConstPORT)

}

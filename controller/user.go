package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/models"
	"github.com/gondsuryaprakash/gondpariwar/utilities"
	"github.com/spf13/cast"
)

type LoginStruct struct {
	Email    string
	Password string
}

func PostLogin(ctx *gin.Context) {
	funcName := "controller.PostLogin"
	logger.I(funcName)
	// v := &LoginStruct{}
	// if err := ctx.Bind(v); err != nil {
	// 	response := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went wrong")
	// 	ctx.JSON(cast.ToInt(utilities.GP_CODE_500), response)
	// 	return
	// }

	// // CheckUser exist
	// isUserExist := models.IsUserExistByEmail(v.Email)
	// if isUserExist {

	// }

}

// Get the user bu user Id - Done
func GetUserById(ctx *gin.Context) {
	funcName := "controller.GetUserById"
	logger.I(funcName)
	userId := ctx.Query("id")
	v, err := models.GetUserById(cast.ToInt(userId))
	if err != nil {
		response := utilities.ResponseWithError("404", "User Data is not available")
		logger.D(err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	returnJson := utilities.ResponseWithModel("200", v, "User Data fetched Successfully", "")
	ctx.IndentedJSON(http.StatusOK, returnJson)

}

// AddUser with user struct Done
func AddUser(ctx *gin.Context) {
	funcName := "controller.user"
	logger.I(funcName)
	var user *models.GpUser
	if err := ctx.Bind(&user); err != nil {
		logger.E(funcName, err)
		resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went wrong")
		ctx.JSON(http.StatusInternalServerError, resturnResponse)
		return
	}

	// Check for use already exist
	isUserExist := models.IsUserExistByEmail(user.Email) // Will Shifted in MiddleWare.
	logger.I("isUserExist", isUserExist)
	if isUserExist {
		response := utilities.ResponseWithError(utilities.GP_CODE_409, "User Already Exist Please login")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_409), response)
		return
	}

	err := models.AddUser(user)
	if err != nil {
		logger.E(funcName, err)
		resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went wrong")
		ctx.JSON(http.StatusInternalServerError, resturnResponse)
		return
	}
	resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_200, "User added successfully")
	ctx.JSON(http.StatusOK, resturnResponse)

}

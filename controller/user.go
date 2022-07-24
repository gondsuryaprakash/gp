package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/models"
	service "github.com/gondsuryaprakash/gondpariwar/service/login"
	"github.com/gondsuryaprakash/gondpariwar/utils"

	"github.com/gondsuryaprakash/gondpariwar/utilities"
	"github.com/spf13/cast"
)

type LoginStruct struct {
	Password string `orm:"column(password);null" json:"password"`
	Email    string `orm:"column(email);null" json:"email"`
}

type GpUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	Religion   string `json:"religion"`
	Dob        string `json:"dob"`
	FatherName string `json:"fathername"`
	MotherName string `json:"mothername"`
}

func PostLogin(ctx *gin.Context) {
	funcName := "controller.PostLogin"
	logger.I(funcName)
	var v *LoginStruct
	if err := ctx.Bind(&v); err != nil {
		logger.E(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := models.GetUserByEmailId(v.Email)

	if err != nil {
		response := utilities.ResponseWithError(utilities.GP_CODE_409, "User doesn't Exist")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_500), response)

	}

	logger.D("user.Password", user.Password)
	logger.D("v.Password", v.Password)
	isPasswordMatched := utils.CheckPasswordHash(user.Password, v.Password)

	if isPasswordMatched {
		token := service.JWTAuthService().GenerateToken(ctx, user.Email)
		logger.I(funcName, token)
		response := utilities.ResponseWithModel(utilities.GP_CODE_200, token, "Loggin Successfully", "")
		ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
		ctx.JSON(cast.ToInt(utilities.GP_CODE_200), response)

	} else {
		response := utilities.ResponseWithError(utilities.GP_CODE_401, "UnAuthorized")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_401), response)
		return
	}
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
	var user *GpUser
	if err := ctx.Bind(&user); err != nil {
		logger.E(funcName, err)
		resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went wrong")
		ctx.JSON(http.StatusInternalServerError, resturnResponse)
		return
	}
	logger.D("user.Password", user.Password)
	user.Password, _ = utils.HashPassword(user.Password)
	// Check for use already exist
	isUserExist, _ := models.IsUserExistByEmail(user.Email) // Will Shifted in MiddleWare.
	logger.I("isUserExist", isUserExist)
	if isUserExist {
		response := utilities.ResponseWithError(utilities.GP_CODE_409, "User Already Exist Please login")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_409), response)
		return
	}
	modelUser := &models.GpUser{
		Name:       user.Name,
		Password:   user.Password,
		Email:      user.Email,
		Mobile:     user.Mobile,
		Gender:     user.Gender,
		Age:        user.Age,
		Religion:   user.Religion,
		Dob:        user.Dob,
		FatherName: user.FatherName,
		MotherName: user.MotherName,
	}
	err := models.AddUser(modelUser)
	if err != nil {
		logger.E(funcName, err)
		resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went wrong")
		ctx.JSON(http.StatusInternalServerError, resturnResponse)
		return
	}
	resturnResponse := utilities.ResponseWithError(utilities.GP_CODE_200, "User added successfully")
	ctx.JSON(http.StatusOK, resturnResponse)

}

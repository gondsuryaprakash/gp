package loginservice

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/models"
	"github.com/gondsuryaprakash/gondpariwar/utilities"
	"github.com/spf13/cast"
)

type ForgotPassword struct {
	Email       string `json:"email"`
	Newpassword string `json:"newpassword"`
}

func Authorise() gin.HandlerFunc {
	funcName := "loginserviceAuthorise"
	logger.I(funcName)
	return Authorization

}

func Authorization(ctx *gin.Context) {

	funcName := "Service.Authorization"
	logger.I(funcName)

	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		response := utilities.ResponseWithError(utilities.GP_CODE_401, "Unautorized")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_401), response)
		ctx.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := JWTAuthService().VerifyJWt(tokenString)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logger.I(claims)
		ctx.Next()
	} else {
		logger.E(funcName, err)
		response := utilities.ResponseWithError(utilities.GP_CODE_401, "Unautorized")
		ctx.JSON(cast.ToInt(utilities.GP_CODE_401), response)
		ctx.Abort()
		return
	}
}

func IsUserExist() gin.HandlerFunc {

	funcName := "loginserviceIsUserExist"
	logger.I(funcName)

	return func(ctx *gin.Context) {
		var v *ForgotPassword
		if err := ctx.Bind(&v); err != nil {
			logger.E(funcName, err)
			response := utilities.ResponseWithError(utilities.GP_CODE_500, "Something went worng")
			ctx.JSON(cast.ToInt(utilities.GP_CODE_500), response)
			ctx.Abort()
			return
		}

		logger.D(funcName)
		existingUser, err := models.GetUserByEmailId(v.Email)
		if err != nil {
			response := utilities.ResponseWithError(utilities.GP_CODE_409, "Email doesn't Exist")
			ctx.JSON(cast.ToInt(utilities.GP_CODE_409), response)
			ctx.Abort()
			return
		}

		ctx.Set("userId", existingUser.Id)
		ctx.Set("newPassword", v.Newpassword)
		ctx.Next()
	}
}

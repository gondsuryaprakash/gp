package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/utilities"
	"github.com/spf13/cast"
)

func Authorise() gin.HandlerFunc {
	funcName := "service.Authorise"
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

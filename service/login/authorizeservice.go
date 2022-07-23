package service

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/logger"
)

func Authorise() gin.HandlerFunc {
	return Authorization

}

func Authorization(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := JWTAuthService().VerifyJWt(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		logger.I(claims)
	} else {
		logger.E(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}

package loginservice

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gondsuryaprakash/gondpariwar/logger"
	"github.com/gondsuryaprakash/gondpariwar/utils"
)

var jwtSecreteKey = utils.GetConfigValue("SECRETKEY")

type AuthService interface {
	VerifyJWt(token string) (*jwt.Token, error)
	GenerateToken(ctx *gin.Context, email string) string
	// setTokenInCookie() error
}

type AuthStruct struct {
	SecretKey string
	Issuer    string
}

type Claims struct {
	Email string
	jwt.StandardClaims
}

func JWTAuthService() AuthService {
	funcName := "loginserviceJWTAuthService"
	logger.D(jwtSecreteKey)
	logger.D(funcName)
	return &AuthStruct{
		SecretKey: jwtSecreteKey,
		Issuer:    "gp_auth",
	}
}

func (v *AuthStruct) VerifyJWt(encodedToken string) (*jwt.Token, error) {
	funcName := "loginserviceVerifyJWt"
	logger.I(funcName)
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			logger.D(funcName, token.Header["alg"])
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(v.SecretKey), nil
	})
}

func (v *AuthStruct) GenerateToken(ctx *gin.Context, email string) string {
	funcName := "loginserviceGenerateToken"
	logger.D(funcName)
	expiryTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
			Issuer:    v.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	tokenString, err := token.SignedString([]byte(v.SecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// func (v *AuthStruct) setTokenInCookie(token string, expireTime time.Time) error {
// 	http.SetCookie(, &http.Cookie{
// 		Name:    "token",
// 		Value:   token,
// 		Expires: expireTime,
// 	})
// }

package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gondsuryaprakash/gondpariwar/utils"
)

var jwtSecreteKey = utils.GetConfigValue("SECRETKEY")

type AuthService interface {
	VerifyJWt(token string) (*jwt.Token, error)
	GenerateToken(email string) string
	// setTokenInCookie() error
}

type AuthStruct struct {
	SecretKey string
}

type Claims struct {
	Email string
	jwt.StandardClaims
}

func JWTAuthService() AuthService {
	return &AuthStruct{
		SecretKey: jwtSecreteKey,
	}
}

func (v *AuthStruct) VerifyJWt(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(jwtSecreteKey), nil
	})
}

func (v *AuthStruct) GenerateToken(email string) string {

	expiryTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, _ := token.SignedString([]byte(v.SecretKey))
	return tokenString
}

// func (v *AuthStruct) setTokenInCookie(token string, expireTime time.Time) error {
// 	http.SetCookie(, &http.Cookie{
// 		Name:    "token",
// 		Value:   token,
// 		Expires: expireTime,
// 	})
// }

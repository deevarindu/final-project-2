package jwt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = "secret"

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := parseToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errReponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errReponse
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errReponse
		}
		return []byte(secret), nil
	})
	fmt.Println(token)
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errReponse
	}
	fmt.Println(token.Claims.(jwt.MapClaims))
	return token.Claims.(jwt.MapClaims), nil
}

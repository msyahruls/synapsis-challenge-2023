package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Level string `json:"level"`
	jwt.StandardClaims
}

var SecretKey = os.Getenv("SECRET_KEY")

func GenerateJwt(issuer, level string) (string, error) {
	session, _ := strconv.Atoi(os.Getenv("SESSION_LOGIN"))
	claims := &Claims{
		Level: level,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(session)).Unix(),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokens.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, string, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", "", err
	}

	return claims.Issuer, claims.Level, nil
}

package utils

import (
	"errors"
	"os"
	"time"

	"github.com/NavenduDuari/go-echo-blog/src/api/model"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey []byte = []byte(os.Getenv("YO_AUTH_TOKEN_SECRET"))

type LoginClaims struct {
	model.UserJWT
	jwt.StandardClaims
}

func GetJwtLoginToken(user model.UserJWT) (string, error) {

	claims := &LoginClaims{
		UserJWT: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJwtLoginToken(tokenString string) (model.UserJWT, error) {
	claims := &LoginClaims{}
	user := model.UserJWT{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtKey, nil
	})
	if err != nil {
		return user, errors.New("Invalid Token")
	}

	if token.Valid {
		return claims.UserJWT, nil
	}
	return user, errors.New("Invalid Token")
}

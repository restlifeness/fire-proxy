package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/restlifeness/fire-proxy.git/src/api/schemas"
)

func getJWTSecretFromEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv("JWT_SECRET")
}

// GenerateJWTToken generates a new JWT token for the given user.
func GenerateJWTToken(UserAuth schemas.RequestAuthForm) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = UserAuth.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(getJWTSecretFromEnv()))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// DecodeJWTToken decodes a given JWT token string and returns the user data.
func DecodeJWTToken(tokenString string) (*schemas.RequestAuthForm, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return getJWTSecretFromEnv(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return nil, errors.New("unable to parse username from JWT")
		}
		return &schemas.RequestAuthForm{
			Username: username,
		}, nil
	}

	return nil, errors.New("invalid JWT token")
}

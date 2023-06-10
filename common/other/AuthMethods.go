package other

import (
	"log"
	"os"
	"strconv"
	"time"
	"workspace/model"

	"github.com/golang-jwt/jwt"
)

// AuthMethods - the interface for the authentication methods
func GenerateToken(user model.UserModel) (string, *jwt.Token) {
	token_lifespan, err1 := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err1 != nil {
		log.Fatal(err1)
		return "", nil
	}

	// map claims for the token
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.Id
	// the expiration should be passed from hours to machine time
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

	// generate the token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// get the secret key from the environment variable
	secretTokenKey := []byte(os.Getenv("JWT_SECRET"))

	// generate token string or return error, please note that the secret should be in the environment variable
	tokenString, err2 := token.SignedString(secretTokenKey)

	if err2 != nil {
		log.Fatal(err2)
		return "", nil
	}

	return tokenString, token
}

// verify access token method
func VerifyAccessToken(tokenString string) (jwt.MapClaims, bool) {
	// get the secret key from the environment variable
	secretTokenKey := []byte(os.Getenv("JWT_SECRET"))

	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretTokenKey, nil
	})

	// check if error
	if err != nil {
		return nil, false
	}

	// check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// return the claims
		return token.Claims.(jwt.MapClaims), true
	}

	return nil, false
}

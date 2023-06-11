package other

import (
	"log"
	"os"
	"strconv"
	"time"
	"workspace/model"

	"github.com/gin-gonic/gin"
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

// verify access token method and return token claims
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

// validate if user is authenticated return true if authenticated and return user id
func ValidateUserAuthenticatedFromCookies(c *gin.Context) (bool, int) {

	// http request context

	// get token string from secure cookie
	tokenBytes, err := SecureCookieObj.GetValue(nil, c.Request)

	// check if error on getting the token string
	if err != nil {
		return false, 0
	} else {
		// verify the token
		tokenClaims, isTokenValid := VerifyAccessToken(string(tokenBytes))
		if !isTokenValid {
			return false, 0
		} else {
			// get the user id from the token claims
			userId := int(tokenClaims["user_id"].(float64))

			// check if the user id is valid
			if userId <= 0 {
				return false, 0
			} else {
				return true, userId
			}
		}
	}
}

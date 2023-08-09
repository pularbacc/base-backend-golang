package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func main() {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 123
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set token expiration time

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		// Handle error
	}

	// tokenString now contains the signed JWT

	// Parse the token from the string
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		// Handle error
	}

	// Check if the token is valid
	if token.Valid {
		// Token is valid
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64) // Type assertion to get the user ID

		fmt.Println(userID)

		// Use the user ID or other claims as needed
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			// Token is malformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is expired or not yet valid
		} else {
			// Token is not valid for some other reason
		}
	} else {
		// Token is not valid for an unknown reason
	}

}

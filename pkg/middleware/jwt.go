// middleware/jwt.go
package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/disaster_management_backend/config"
)

// Function to generate a JWT token for a user
func GenerateJWT(userID uint, role string, isVerified bool) (string, error) {
    accountVerified := "unverified"
    if isVerified {
        accountVerified = "verified"
    }
	claims := jwt.MapClaims{
		"sub":         userID,
		"role":        role,
		"is_verified": accountVerified,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

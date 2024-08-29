package tokens

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	jwt.RegisteredClaims
}

package utils

import (
	"fmt"
	tokens "hst_manag/internal/app/models/JWT"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string, phone string) (string, error) {
	claims := &tokens.CustomClaims{
		Email:     email,
		FirstName: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"firstname": firstname,
	// 	"email":     email,
	// 	"exp":       time.Now().Add(time.Hour * 24).Unix(),
	// })
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Println("============Generated Token:=1========>>", tokenString)
	return tokenString, nil

}
func ValidateToken(tokenString string) (*tokens.CustomClaims, error) {
	fmt.Println("===============Validating Token:====2=========>>", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &tokens.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error parsing token:", err)
		return nil, err

	}
	claims, ok := token.Claims.(*tokens.CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

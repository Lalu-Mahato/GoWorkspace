package utils

import (
	"github/Lalu-Mahato/GoWorkspace/gin_postgres_boilerplate/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string, id uint64) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &models.JWTClaim{
		Email: email,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	return
}

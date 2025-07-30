package utility

import (
	"cobaaja/contact-management/app/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(user *entity.User) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // Token berlaku 2 jam
	})
	return token.SignedString(jwtSecret)
}

package dto

import "github.com/golang-jwt/jwt/v5"

type JwtLocalsClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

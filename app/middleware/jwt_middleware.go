package middleware

import (
	"cobaaja/contact-management/app/dto"
	"cobaaja/contact-management/utility"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware untuk memverifikasi JWT
func JWTMiddleware(ctx *fiber.Ctx) error {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return utility.UnauthenticatedResponse("Missing or Mailformed JWT", nil, ctx)
	}

	// Format Authorization: Bearer <token>
	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return utility.UnauthenticatedResponse("Invalid Authroization Header", nil, ctx)
	}

	tokenStr := authHeader[7:]

	// Parse dan verifikasi token
	token, err := jwt.ParseWithClaims(tokenStr, &dto.JwtLocalsClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return utility.UnauthenticatedResponse("Invalid or Expired Token JWT", nil, ctx)
	}

	// Validasi token dan claims
	if claims, ok := token.Claims.(*dto.JwtLocalsClaims); ok && token.Valid {
		// Store user information in context for use in subsequent handlers
		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("username", claims.Username)

		// lanjuutt
		return ctx.Next()
	}

	return utility.UnauthenticatedResponse("Invalid Token", nil, ctx)
}

package utility

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterStrongPasswordValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String() // inputan jsonnya

	min8 := len(value) >= 8
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(value)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(value)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(value)

	return min8 && hasUpper && hasLower && hasNumber
}

// func RegisterUniqueUsernameValidation(repo authrepository.AuthRepository) validator.Func {
// 	return func(fl validator.FieldLevel) bool {
// 		value := fl.Field().String()

// 		repo.db.
// 	}
// }

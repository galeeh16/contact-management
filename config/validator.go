package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// buat message validator (sesuai abjad A-Z)
const (
	Unknown = "Unknown message for Tag '%s'" // default unhandled message

	Aplha               = "The %s field must only contain letters."
	AlphaNum            = "The %s field must only contain letters and numbers."
	Boolean             = "The %s field must be true or false."
	Datetime            = "The %s field must match the format %s."
	Eqfield             = "%s must be same with %s."
	GreatherThan        = "The %s field must be greater than %s."
	GreatherThanOrEqual = "The %s field must be greater than or equal to %s."
	Length              = "The %s field must be %s digits."
	LessThan            = "The %s field must be less than %s."
	LessThanOrEqual     = "The %s field must be less than or equal to %s."
	Max                 = "The %s field must not be greater than %s characters."
	MimeType            = "%s must be: %s."
	Min                 = "The %s field must be at least %s characters."
	Number              = "The %s must be a valid number."
	Numeric             = "The %s must be numeric."
	OneOf               = "The selected %s is invalid."
	Required            = "The %s field is required."
	StrongPassword      = "%s minimal 8 karakter, harus mengandung huruf besar, huruf kecil, angka."
	Unique              = "%s has been taken."
)

// Validator adalah wrapper untuk validator instance
type Validator struct {
	Validate *validator.Validate
}

// NewValidator membuat instance validator baru dengan json format
func NewValidator() *Validator {
	v := validator.New()
	return &Validator{Validate: v}
}

// RegisterTagJSON mendaftarkan fungsi untuk menggunakan nama field dari tag json
func (v *Validator) RegisterTagJSON() {
	v.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("form")
		if name == "" {
			name = fld.Name
		}
		// Menghapus ",omitempty" jika ada di tag json
		return strings.SplitN(name, ",", 2)[0]
	})
}

// RegisterTagForm mendaftarkan fungsi untuk menggunakan nama field dari tag form
func (v *Validator) RegisterTagForm() {
	v.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("form")
		if name == "" {
			name = fld.Name
		}
		// Menghapus ",omitempty" jika ada di tag form
		return strings.SplitN(name, ",", 2)[0]
	})
}

// ValidateStruct memvalidasi struct dan mengembalikan ErrorResponse jika ada error
func (v *Validator) ValidateStruct(data any) map[string]string {
	if err := v.Validate.Struct(data); err != nil {
		arrayError := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			var message string

			switch err.Tag() {
			case "alpha":
				message = fmt.Sprintf(Aplha, field)
			case "alphanum":
				message = fmt.Sprintf(AlphaNum, field)
			case "boolean":
				message = fmt.Sprintf(Boolean, field)
			case "datetime":
				message = fmt.Sprintf(Datetime, field, err.Param())
			case "eqfield":
				message = fmt.Sprintf(Eqfield, field, err.Param())
			case "gt":
				message = fmt.Sprintf(GreatherThan, field, err.Param())
			case "gte":
				message = fmt.Sprintf(GreatherThanOrEqual, field, err.Param())
			case "length":
				message = fmt.Sprintf(Length, field, err.Param())
			case "lt":
				message = fmt.Sprintf(LessThan, field, err.Param())
			case "lte":
				message = fmt.Sprintf(LessThanOrEqual, field, err.Param())
			case "max":
				message = fmt.Sprintf(Max, err.Tag(), err.Param())
			case "mimetype":
				message = fmt.Sprintf(MimeType, field, err.Param())
			case "min":
				message = fmt.Sprintf(Min, err.Tag(), err.Param())
			case "number":
				message = fmt.Sprintf(Number, field)
			case "numeric":
				message = fmt.Sprintf(Numeric, field)
			case "oneof":
				message = fmt.Sprintf(OneOf, field)
			case "required":
				message = fmt.Sprintf(Required, field)
			case "strong_password":
				message = fmt.Sprintf(StrongPassword, field)
			case "unique_contact_name":
				message = fmt.Sprintf(Unique, err.Value())
			case "unique_contact_phone":
				message = fmt.Sprintf(Unique, err.Value())
			case "unique_contact_phone_edit":
				message = fmt.Sprintf(Unique, err.Value())
			case "unique_email":
				message = fmt.Sprintf(Unique, err.Value())
			case "unique_username":
				message = fmt.Sprintf(Unique, err.Value())

			default:
				message = fmt.Sprintf(Unknown, err.Tag())
			}

			// errs[err.Field()] = err.Translate(v.validate.Translator)
			arrayError[field] = message
		}
		// return &ErrorResponse{Errors: errs}
		return arrayError
	}
	return nil
}

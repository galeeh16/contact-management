package utility

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator adalah wrapper untuk validator instance
type Validator struct {
	Validate *validator.Validate
}

// NewValidator membuat instance validator baru dengan json format
func NewValidator() *Validator {
	v := validator.New()

	// Gunakan nama field dari tag json
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "" {
			name = fld.Name
		}
		// Menghapus ",omitempty" jika ada di tag json
		return strings.SplitN(name, ",", 2)[0]
	})

	return &Validator{Validate: v}
}

func NewValidatorFormData() *Validator {
	v := validator.New()

	// Gunakan nama field dari tag json
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("form")
		// fmt.Println("name form", name)
		if name == "" {
			name = fld.Name
		}
		// Menghapus ",omitempty" jika ada di tag json
		return strings.SplitN(name, ",", 2)[0]
	})

	return &Validator{Validate: v}
}

// ValidateStruct memvalidasi struct dan mengembalikan ErrorResponse jika ada error
func (v *Validator) ValidateStruct(data any) map[string]string {
	if err := v.Validate.Struct(data); err != nil {
		errs := make(map[string]string)

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
			case "unique_email":
				message = fmt.Sprintf(Unique, err.Value())
			case "unique_username":
				message = fmt.Sprintf(Unique, err.Value())

			default:
				message = fmt.Sprintf(Unknown, err.Tag())
			}

			// errs[err.Field()] = err.Translate(v.validate.Translator)
			errs[field] = message
		}
		// return &ErrorResponse{Errors: errs}
		return errs
	}
	return nil
}

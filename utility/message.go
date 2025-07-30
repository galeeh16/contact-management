package utility

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

package dto

import "time"

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=50,unique_username"`
	Name            string `json:"name" validate:"required,min=3,max=100"`
	Password        string `json:"password" validate:"required,max=50,strong_password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type RegisterResponse struct {
	ID        uint64     `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

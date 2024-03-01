package model

import "time"

type UserRegister struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Avatar   string `json:"avatar"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" `
	FullName  string    `json:"full_name" `
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

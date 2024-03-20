package model

import (
	"github.com/google/uuid"
	"mime/multipart"
)

type UserRegister struct {
	Email           string `json:"email" binding:"required,email"`
	FullName        string `json:"full_name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email" `
	FullName string    `json:"full_name" `
	Avatar   string    `json:"avatar"`
}

type UserUpdate struct {
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
	PhotoLink  string `json:"photo_link"`
}

type UserUploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}

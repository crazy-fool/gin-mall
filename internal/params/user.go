package params

import "time"

type RegisterParam struct {
	Name        *string `json:"name" binding:"required"`
	Password    *string `json:"password" binding:"required"`
	Account     string  `json:"account" binding:"required"`
	Mobile      *string `json:"mobile" binding:"required"`
	LastLoginAt *time.Time
	Id          int32 `json:"id"`
}

type LoginParam struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileParam struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordParam struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

// user login, user tokens, user reset password
package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginLog struct {
	gorm.Model
	UserID    uint   `json:"user_id" validate:"required"`
	IPAddress string `gorm:"size:45" json:"ip_address" validate:"required"` // support ipv4, ipv6
	Device    string `gorm:"size:100" json:"device" validate:"required"`
}

type UserToken struct {
	gorm.Model
	UserID     uint      `json:"user_id" validate:"required"`
	Token      string    `gorm:"uniqueIndex;size:255" json:"token" validate:"required"`
	ExpireTime time.Time `json:"expire_time" validate:"required"` // Token expire time
}

type PasswordReset struct {
	gorm.Model
	Email      string    `gorm:"size:100" json:"email" validate:"omitempty"`
	Phone      string    `gorm:"size:20" json:"phone" validate:"omitempty"`
	Token      string    `gorm:"size:255" json:"token" validate:"required"`
	ExpireTime time.Time `json:"expire_time" validate:"required"` // Token expire time
}

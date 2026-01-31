package model

import (
	"gorm.io/gorm"
)

type UserStatus string

const (
	StatusActive   UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBanned   UserStatus = "BANNED"
)

type UserRole string

const (
	RoleUser      UserRole = "user"
	RoleAdmin     UserRole = "admin"
	RoleModerator UserRole = "moderator"
)

type AuthModel struct {
	gorm.Model
	Username string     `json:"username" gorm:"type:varchar(100);not null;uniqueIndex"`
	Password string     `json:"password" gorm:"type:varchar(100);not null"`
	Email    string     `json:"email" gorm:"type:varchar(100);not null;uniqueIndex"`
	Phone    string     `json:"phone" gorm:"type:varchar(15);not null;uniqueIndex"`
	Status   UserStatus `json:"status" gorm:"type:varchar(20);default:'INACTIVE'"`
	Role     UserRole   `json:"role" gorm:"type:varchar(50);not null;default:'user'"`
}

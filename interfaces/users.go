package interfaces

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

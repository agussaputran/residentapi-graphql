package tables

import "gorm.io/gorm"

// Users table
type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

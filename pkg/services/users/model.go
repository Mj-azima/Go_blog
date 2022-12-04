package users

import "gorm.io/gorm"

//User model
type Users struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"password"`
}

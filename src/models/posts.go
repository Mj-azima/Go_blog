package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Body   string `gorm:"size:10000"`
	Author uint   `json:"author"`
}

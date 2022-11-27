package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Body     string `gorm:"size:10000" json:"body"`
	AuthorID uint   `json:"authorID"`
	Author   Users  `json:"author"`
}

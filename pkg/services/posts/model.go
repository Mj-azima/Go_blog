package posts

import (
	"blog/pkg/services/users"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Body     string        `gorm:"size:10000" json:"body"`
	AuthorID int           `json:"authorID"`
	Author   users.Users   `json:"author"`
	Editors  []users.Users `gorm:"many2many:post_editors;" json:"editors"`
}

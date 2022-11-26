package interfaces

import (
	"blog/src/models"
	"gorm.io/gorm"
)

type IPost interface {
	Get(i int) (models.Posts, *gorm.DB, error)
	GetByIdAndAuthor(user, post int) (models.Posts, error)
	GetAll() (models.Posts, error)
	Create(author models.Users, body string) error
	Edit(id, author uint, body string) error
	Delete(id int) error
}

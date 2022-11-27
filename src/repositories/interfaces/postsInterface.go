package interfaces

import (
	"blog/src/models"
	"gorm.io/gorm"
)

type IPost interface {
	Get(id int) (models.Posts, *gorm.DB, error)
	GetByIdAndAuthor(userId uint, id int) (models.Posts, error)
	GetAll() ([]models.Posts, error)
	Create(author models.Users, body string) error
	Edit(id int, body string) error
	Delete(id int) error
}

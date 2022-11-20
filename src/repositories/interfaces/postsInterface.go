package interfaces

import (
	"blog/src/models"
)

type IPost interface {
	Get(i int) (models.Posts, error)
	GetAll() (models.Posts, error)
	Create(author uint, body string) error
	Edit(id, author uint, body string) error
	Delete(id int) error
}

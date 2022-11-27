package interfaces

import "blog/src/models"

type IPost interface {
	Create(user, body string) error
	Update(id int, body string) error
	Get(id int) (models.Posts, models.Users, error)
	GetAll() ([]models.Posts, error)
	Delete(id int) error
}

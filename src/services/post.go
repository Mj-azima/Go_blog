package services

import (
	"blog/src/database"
	"blog/src/models"
	"blog/src/repositories"
	"log"
)

var userModel *repositories.User

var postModel *repositories.Post

func init() {
	postModel = new(repositories.Post)
	userModel = new(repositories.User)

}

type Post struct {
}

func (p *Post) Create(email, body string) error {

	user, err := userModel.GetByEmail(email)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		log.Fatal("not found a user")
	}

	if err := postModel.Create(user, body); err != nil {
		return err
	}

	return nil
}

func (p *Post) Update(id int, body string) error {

	if err := postModel.Edit(id, body); err != nil {
		return err
	}
	return nil
}

func (p *Post) Get(id int) (models.Posts, models.Users, error) {
	post, _, err := postModel.Get(id)
	if err != nil {
		return post, models.Users{}, err
	}
	var user models.Users

	if err := database.DBConn.First(&user, post.Author).Error; err != nil {
		return models.Posts{}, user, err
	}
	return post, user, nil
}

func (p *Post) GetAll() ([]models.Posts, error) {
	posts, err := postModel.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Post) Delete(id int) error {
	if err := postModel.Delete(id); err != nil {
		return err
	}
	return nil
}

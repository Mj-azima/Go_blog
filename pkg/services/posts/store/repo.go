package store

import (
	"blog/pkg/services/posts"
	"blog/pkg/services/users"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type postRepo struct {
	DB *gorm.DB
}

var once sync.Once
var singleInstance *postRepo

func New(conn *gorm.DB) posts.Repo {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &postRepo{conn}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance

	//return &postRepo{conn}
}

func (p *postRepo) Get(id int) (posts.Posts, *gorm.DB, error) {
	var post posts.Posts
	result := p.DB.First(&post, id)
	if err := result.Error; err != nil {
		return post, result, err
	}

	return post, result, nil
}

func (p *postRepo) GetByIdAndAuthor(userId uint, id int) (posts.Posts, error) {
	var post posts.Posts
	if err := p.DB.Find(&post, "author_id = ? AND id = ?", userId, id).Error; err != nil {
		return post, err
	}
	return post, nil

}

func (p *postRepo) GetAll() ([]posts.Posts, error) {
	var allpost []posts.Posts

	result := p.DB.Preload("Author").Find(&allpost)

	if result.Error != nil {
		return allpost, result.Error
	}
	return allpost, nil
}

func (p *postRepo) Create(author users.Users, body string) error {
	post := posts.Posts{
		Author: author,
		Body:   body,
	}
	tx := p.DB.Create(&post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *postRepo) Edit(id int, body string) error {

	post, _, err := p.Get(id)
	if err != nil {
		return err
	}

	result := p.DB.Model(&post).Update("Body", body)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *postRepo) Delete(id int) error {
	var post posts.Posts
	if err := p.DB.Delete(&post, id).Error; err != nil {
		return err
	}
	return nil
}

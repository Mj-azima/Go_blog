package posts

import (
	"blog/pkg/services/users"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type Repo interface {
	Get(id int) (Posts, *gorm.DB, error)
	GetByIdAndAuthor(userId uint, id int) (Posts, error)
	GetAll() ([]Posts, error)
	Create(author users.Users, body string) error
	Edit(id int, body string, user users.Users) error
	Delete(id int) error
}

type Service interface {
	Create(author users.Users, body string) error
	Update(id int, body string, user users.Users) error
	Get(id int) (Posts, error)
	GetAll() ([]Posts, error)
	Delete(id int) error
}

type post struct {
	repo Repo
}

var once sync.Once
var singleInstance *post

func New(repo Repo) Service {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &post{repo}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
	//return &post{repo}
}

func (p *post) Create(author users.Users, body string) error {
	if err := p.repo.Create(author, body); err != nil {
		return err
	}
	return nil
}

func (p *post) Update(id int, body string, user users.Users) error {

	if err := p.repo.Edit(id, body, user); err != nil {
		return err
	}
	return nil
}

func (p *post) Get(id int) (Posts, error) {
	post, _, err := p.repo.Get(id)
	if err != nil {
		return post, err
	}
	//var user users.Users
	//
	//if err := database.DBConn.First(&user, post.Author).Error; err != nil {
	//	return Posts{}, err
	//}
	return post, nil
}

func (p *post) GetAll() ([]Posts, error) {
	posts, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *post) Delete(id int) error {
	if err := p.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

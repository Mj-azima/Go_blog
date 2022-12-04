package posts

import (
	"blog/pkg/services/users"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

//Repository interface
type Repo interface {
	Get(id int) (Posts, *gorm.DB, error)
	GetByIdAndAuthor(userId uint, id int) (Posts, error)
	GetAll() ([]Posts, error)
	Create(author users.Users, body string) error
	Edit(id int, body string, user users.Users) error
	Delete(id int) error
}

//Service interface
type Service interface {
	Create(author users.Users, body string) error
	Update(id int, body string, user users.Users) error
	Get(id int) (Posts, error)
	GetAll() ([]Posts, error)
	Delete(id int) error
}

//Post struct
type post struct {
	repo Repo
}

//Constructor with singleton pattern
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

//Create method service
func (p *post) Create(author users.Users, body string) error {
	//Create post in repository
	if err := p.repo.Create(author, body); err != nil {
		return err
	}
	return nil
}

//Update method service
func (p *post) Update(id int, body string, user users.Users) error {
	//Update Post in repository
	if err := p.repo.Edit(id, body, user); err != nil {
		return err
	}
	return nil
}

//Get method service
func (p *post) Get(id int) (Posts, error) {
	//Get post from repository
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

//Get all method service
func (p *post) GetAll() ([]Posts, error) {
	//Get all post from repository
	posts, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

//Delete method service
func (p *post) Delete(id int) error {
	//Delete post from repository
	if err := p.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

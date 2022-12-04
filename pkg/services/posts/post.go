package posts

import (
	"blog/pkg/services/users"
	"fmt"
	"sync"
)

//Repository interface
type Repo interface {
	Get(id int) (Posts, error)
	GetByIdAndAuthor(userId uint, id int) (Posts, error)
	GetAll() ([]Posts, error)
	Create(author users.Users, body string) (Posts, error)
	Edit(id int, body string, user users.Users) (Posts, error)
	Delete(id int) (Posts, error)
}

//Service interface
type Service interface {
	Create(author users.Users, body string) (Posts, error)
	Update(id int, body string, user users.Users) (Posts, error)
	Get(id int) (Posts, error)
	GetAll() ([]Posts, error)
	Delete(id int) (Posts, error)
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
func (p *post) Create(author users.Users, body string) (Posts, error) {
	//Create post in repository
	post, err := p.repo.Create(author, body)
	if err != nil {
		return post, err
	}
	return post, nil
}

//Update method service
func (p *post) Update(id int, body string, user users.Users) (Posts, error) {
	//Update Post in repository
	post, err := p.repo.Edit(id, body, user)
	if err != nil {
		return post, err
	}
	return post, nil
}

//Get method service
func (p *post) Get(id int) (Posts, error) {
	//Get post from repository
	post, err := p.repo.Get(id)
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
func (p *post) Delete(id int) (Posts, error) {
	//Delete post from repository
	post, err := p.repo.Delete(id)
	if err != nil {
		return post, err
	}
	return post, nil
}

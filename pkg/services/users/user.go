package users

import (
	"blog/pkg/utils/sessions"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sync"
)

//Repository interface
type Repo interface {
	Get(id int) (Users, error)
	GetByEmail(email string) (Users, error)
	GetAll() ([]Users, error)
	Create(email string, password []byte) error
	Edit(id int, email string, password []byte) error
	Delete(id int) error
}

//Service interface
type Service interface {
	Create(email string, password []byte) error
	Update(email string, password []byte) error
	Get(id int) (Users, error)
	GetByEmail(email string) (Users, error)
	GetAll() ([]Users, error)
	Delete(id int) error
	IsLogin(c *fiber.Ctx) (bool, error)
}

//User struct
type user struct {
	repo Repo
}

//Constructor with singleton pattern
var once sync.Once
var singleInstance *user

func New(repo Repo) Service {

	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &user{repo}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
	//return &user{repo}
}

//Create method service
func (u *user) Create(email string, password []byte) error {
	//Create user in repository
	if err := u.repo.Create(email, password); err != nil {
		return err
	}
	return nil
}

//Get by email method service
func (u *user) GetByEmail(email string) (Users, error) {
	//Get by email user from repository
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

//Get method service
func (u *user) Get(id int) (Users, error) {
	//Get user from repository
	user, err := u.repo.Get(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (a *user) IsLogin(c *fiber.Ctx) (bool, error) {

	//Get user session
	Session := sessions.Instance
	user, err := Session.Get(c)
	if err != nil {
		return false, err
	}

	//Check user is authenticate
	if user == nil {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return false, nil
	}
	return true, nil
}

//Update method service
func (u *user) Update(email string, password []byte) error {
	//TODO implement me
	panic("implement me")
}

//Delete method service
func (u *user) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

//Get all method service
func (u *user) GetAll() ([]Users, error) {
	//TODO implement me
	panic("implement me")
}

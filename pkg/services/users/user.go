package users

import (
	"blog/pkg/utils/sessions"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Repo interface {
	Get(id int) (Users, error)
	GetByEmail(email string) (Users, error)
	GetAll() ([]Users, error)
	Create(email string, password []byte) error
	Edit(id int, email string, password []byte) error
	Delete(id int) error
}

type Service interface {
	Create(email string, password []byte) error
	Update(email string, password []byte) error
	Get(id int) (Users, error)
	GetByEmail(email string) (Users, error)
	GetAll() ([]Users, error)
	Delete(id int) error
	IsLogin(c *fiber.Ctx) (bool, error)
}

type user struct {
	repo Repo
}

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

func (u *user) Create(email string, password []byte) error {
	if err := u.repo.Create(email, password); err != nil {
		return err
	}
	return nil
}

func (u *user) GetByEmail(email string) (Users, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *user) Update(email string, password []byte) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) Get(id int) (Users, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *user) GetAll() ([]Users, error) {
	//TODO implement me
	panic("implement me")
}

func (a *user) IsLogin(c *fiber.Ctx) (bool, error) {
	Session := sessions.Instance

	user, err := Session.Get(c)
	if err != nil {
		return false, err
	}

	if user == nil {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return false, nil
	}
	return true, nil
}

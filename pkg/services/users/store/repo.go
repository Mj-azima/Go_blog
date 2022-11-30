package store

import (
	"blog/pkg/services/users"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type userRepo struct {
	DB *gorm.DB
}

var once sync.Once
var singleInstance *userRepo

func New(conn *gorm.DB) users.Repo {

	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &userRepo{conn}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance

	//return &userRepo{conn}
}

func (u *userRepo) Get(id int) (users.Users, error) {
	var user users.Users
	if err := u.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepo) GetByEmail(email string) (users.Users, error) {

	var user users.Users

	if err := u.DB.Find(&user, "email = ?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepo) GetAll() ([]users.Users, error) {
	var allUser []users.Users
	result := u.DB.Find(&allUser)
	if result.Error != nil {
		return allUser, result.Error
	}
	return allUser, nil
}

func (u *userRepo) Create(email string, password []byte) error {
	user := users.Users{
		Email:    email,
		Password: password,
	}
	tx := u.DB.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *userRepo) Edit(id int, email string, password []byte) error {

	user, err := u.Get(id)
	if err != nil {
		return err
	}
	result := u.DB.Model(&user).Updates(users.Users{Email: email, Password: password})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepo) Delete(id int) error {
	var user users.Users
	if err := u.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

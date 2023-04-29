package store

import (
	"blog/pkg/services/users"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

//var once sync.Once
//var singleInstance *userRepo

func New(conn *gorm.DB) users.Repo {

	//if singleInstance == nil {
	//	once.Do(
	//		func() {
	//			fmt.Println("Creating single instance now.")
	//			singleInstance = &userRepo{conn}
	//		})
	//} else {
	//	fmt.Println("Single instance already created.")
	//}
	//
	//return singleInstance

	return &userRepo{conn}
}

func (u *userRepo) Get(id int) (users.Users, error) {
	var user users.Users
	if err := u.DB.First(&user, id).Error; err != nil {
		return user, users.ErrUserNotFound
	}

	return user, nil
}

func (u *userRepo) GetByEmail(email string) (users.Users, error) {

	var user users.Users

	if err := u.DB.Find(&user, "email = ?", email).Error; err != nil {

		return user, users.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepo) GetAll() ([]users.Users, error) {
	var allUser []users.Users
	result := u.DB.Find(&allUser)
	if result.Error != nil {
		return allUser, users.ErrUserQuery
	}
	return allUser, nil
}

func (u *userRepo) Create(email string, password []byte) (uint, error) {
	if email == "" {
		return 0, users.ErrUserCreate
	}
	if password == nil {
		return 0, users.ErrUserCreate
	}
	user := users.Users{
		Email:    email,
		Password: password,
	}
	tx := u.DB.Create(&user)
	if tx.Error != nil {
		return user.ID, users.ErrUserCreate
	}
	return user.ID, nil
}

func (u *userRepo) Edit(id int, email string, password []byte) error {

	user, err := u.Get(id)
	if err != nil {
		return users.ErrUserNotFound
	}
	result := u.DB.Model(&user).Updates(users.Users{Email: email, Password: password})
	if result.Error != nil {
		return users.ErrUserUpdate
	}
	return nil
}

func (u *userRepo) Delete(id int) (users.Users, error) {
	var user users.Users
	if err := u.DB.Delete(&user, id).Error; err != nil {
		return user, users.ErrUserDelete
	}
	return user, nil
}

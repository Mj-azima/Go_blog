package repositories

import (
	"blog/src/database"
	"blog/src/models"
)

type User models.Users

func (p *User) Get(id int) (models.Users, error) {
	var user models.Users
	if err := database.DBConn.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (p User) GetAll() ([]models.Users, error) {
	var users []models.Users
	result := database.DBConn.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (p User) Create(email string, password []byte) error {
	user := models.Users{
		Email:    email,
		Password: password,
	}
	tx := database.DBConn.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *User) Edit(id int, email string, password []byte) error {
	//var user models.Users
	//result := database.DBConn.Find(&user, id)
	//if result.Error != nil {
	//	return result.Error
	//}
	user, err := p.Get(id)
	if err != nil {
		return err
	}
	result := database.DBConn.Model(&user).Updates(models.Users{Email: email, Password: password})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *User) Delete(id int) error {
	var user models.Users
	if err := database.DBConn.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

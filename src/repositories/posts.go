package repositories

import (
	"blog/src/database"
	"blog/src/models"
	"gorm.io/gorm"
)

type Post struct {
}

func (p *Post) Get(id int) (models.Posts, *gorm.DB, error) {
	var post models.Posts
	result := database.DBConn.First(&post, id)
	if err := result.Error; err != nil {
		return post, result, err
	}

	return post, result, nil
}

func (p *Post) GetByIdAndAuthor(user uint, id int) (models.Posts, error) {
	var post models.Posts
	if err := database.DBConn.Find(&post, "auther = ? AND id = ?", user, id).Error; err != nil {
		return post, err
	}
	return post, nil

}

func (p *Post) GetAll() ([]models.Posts, error) {
	var posts []models.Posts
	result := database.DBConn.Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}
	return posts, nil
}

func (p *Post) Create(author models.Users, body string) error {
	post := models.Posts{
		Author: author,
		Body:   body,
	}
	tx := database.DBConn.Create(&post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *Post) Edit(id int, body string) error {

	post, _, err := p.Get(id)
	if err != nil {
		return err
	}

	result := database.DBConn.Model(&post).Update("Body", body)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Post) Delete(id int) error {
	var post models.Posts
	if err := database.DBConn.Delete(&post, id).Error; err != nil {
		return err
	}
	return nil
}

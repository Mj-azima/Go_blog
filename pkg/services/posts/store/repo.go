package store

import (
	"blog/pkg/services/posts"
	"blog/pkg/services/users"
	"gorm.io/gorm"
)

//post repository struct
type postRepo struct {
	DB *gorm.DB
}

//repository constructor with singleton pattern
//var once sync.Once
//var singleInstance *postRepo

func New(conn *gorm.DB) posts.Repo {
	//if singleInstance == nil {
	//	once.Do(
	//		func() {
	//			fmt.Println("Creating single instance now.")
	//			singleInstance = &postRepo{conn}
	//		})
	//} else {
	//	fmt.Println("Single instance already created.")
	//}
	//
	//return singleInstance

	return &postRepo{conn}
}

//Get method repository
func (p *postRepo) Get(id int) (posts.Posts, error) {
	var post posts.Posts
	result := p.DB.Preload("Editors").First(&post, id)
	if err := result.Error; err != nil {
		return post, posts.ErrPostNotFound
	}

	return post, nil
}

//Get by id & author method repository
func (p *postRepo) GetByIdAndAuthor(userId uint, id int) (posts.Posts, error) {
	var post posts.Posts
	if err := p.DB.Find(&post, "author_id = ? AND id = ?", userId, id).Error; err != nil {
		return post, posts.ErrPostQuery
	}
	return post, nil

}

//Get all method repository
func (p *postRepo) GetAll() ([]posts.Posts, error) {
	var allpost []posts.Posts

	result := p.DB.Preload("Author").Find(&allpost)

	if result.Error != nil {
		return allpost, posts.ErrPostQuery
	}
	return allpost, nil
}

//Create method repository
func (p *postRepo) Create(authorId int, body string) (uint, error) {
	post := posts.Posts{
		AuthorID: authorId,
		Body:     body,
	}
	tx := p.DB.Create(&post)
	if tx.Error != nil {
		return post.ID, posts.ErrPostCreate
	}
	return post.ID, nil
}

//Update method repository
func (p *postRepo) Edit(id int, body string, user users.Users) error {

	post, err := p.Get(id)
	if err != nil {
		return posts.ErrPostNotFound
	}

	err = p.DB.Model(&post).Update("Body", body).Association("Editors").Append(&user)
	if err != nil {
		return posts.ErrPostUpdate
	}
	return nil
}

//Delete method repository
func (p *postRepo) Delete(id int) (posts.Posts, error) {
	var post posts.Posts
	if err := p.DB.Delete(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}

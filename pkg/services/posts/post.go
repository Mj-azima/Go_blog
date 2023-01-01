package posts

import (
	"blog/pkg/services/users"
)

//Repository interface
type Repo interface {
	Get(id int) (Posts, error)
	GetByIdAndAuthor(userId uint, id int) (Posts, error)
	GetAll() ([]Posts, error)
	Create(authorId int, body string) (uint, error)
	Edit(id int, body string, user users.Users) error
	Delete(id int) (Posts, error)
}

//Service interface
type Service interface {
	Create(authorId int, body string) (Posts, error)
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
//var once sync.Once
//var singleInstance *post

func New(repo Repo) Service {
	//if singleInstance == nil {
	//	once.Do(
	//		func() {
	//			fmt.Println("Creating single instance now.")
	//			singleInstance = &post{repo}
	//		})
	//} else {
	//	fmt.Println("Single instance already created.")
	//}
	//
	//return singleInstance
	return &post{repo}
}

//Create method service
func (p *post) Create(authorId int, body string) (Posts, error) {
	//Create post in repository
	postId, err := p.repo.Create(authorId, body)
	if err != nil {
		return Posts{}, err
	}
	return p.repo.Get(int(postId))
}

//Update method service
func (p *post) Update(id int, body string, user users.Users) (Posts, error) {
	//Update Post in repository
	err := p.repo.Edit(id, body, user)
	if err != nil {
		return Posts{}, err
	}
	return p.repo.Get(id)
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

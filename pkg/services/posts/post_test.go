package posts

import (
	"blog/pkg/services/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type repoMock struct {
	GetResult Posts
	GetError  error

	GetAllResult []Posts
	GetAllError  error

	CreateResult uint
	CreateError  error

	UpdateError error

	DeleteResult Posts
	DeleteError  error
}

func (r *repoMock) Get(id int) (Posts, error) {
	return r.GetResult, r.GetError
}

func (r *repoMock) GetByIdAndAuthor(userId uint, id int) (Posts, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repoMock) GetAll() ([]Posts, error) {
	return r.GetAllResult, r.GetAllError
}

func (r *repoMock) Create(author users.Users, body string) (uint, error) {
	return r.CreateResult, r.CreateError
}

func (r *repoMock) Edit(id int, body string, user users.Users) error {
	return r.UpdateError
}

func (r *repoMock) Delete(id int) (Posts, error) {
	return r.DeleteResult, r.DeleteError
}

func TestServiceGet(t *testing.T) {
	id := uuid.New().ID()

	tests := map[string]struct {
		repo   Repo
		result Posts
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				GetResult: Posts{
					Model: gorm.Model{ID: uint(id)},
				},
				GetError: nil,
			},
			result: Posts{
				Model: gorm.Model{ID: uint(id)},
			},
			err: nil,
		},
		"Not found from repo": {
			repo: &repoMock{
				GetResult: Posts{},
				GetError:  ErrPostNotFound,
			},
			result: Posts{},
			err:    ErrPostNotFound,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := New(test.repo)
			response, err := service.Get(int(id))

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result.ID, response.ID)
		})
	}
}

func TestServiceCreate(t *testing.T) {
	postId := uuid.New().ID()
	//title := "some-title"
	body := "some-body"
	//ac := PostsCreateUpdate{Title: title, Body: body}
	ui := users.Users{
		Email:    "admin@example.com",
		Password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
	}

	pr := Posts{
		Model: gorm.Model{
			ID: uint(postId),
			//CreatedAt: time.Time{},
			//UpdatedAt: time.Time{},
		},
		Author: ui,
		//Title:      title,
		Body: body,
	}

	tests := map[string]struct {
		repo   Repo
		result Posts
		body   string
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				CreateResult: uint(postId),
				CreateError:  nil,
				GetResult:    pr,
				GetError:     nil,
			},
			//input:  ac,
			body:   "some-body",
			result: pr,
			err:    nil,
		},
		"Not found from repo after create": {
			repo: &repoMock{
				CreateResult: uint(postId),
				CreateError:  nil,
				GetResult:    Posts{},
				GetError:     ErrPostNotFound,
			},
			//input:  ac,
			body:   "some-body",
			result: Posts{},
			err:    ErrPostNotFound,
		},
		"Creation failure": {
			repo: &repoMock{
				CreateResult: uint(postId),
				CreateError:  ErrPostCreate,
			},
			//input:  ac,
			body:   "some-body",
			result: Posts{},
			err:    ErrPostCreate,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := New(test.repo)
			post, err := service.Create(ui, test.body)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result, post)

		})
	}
}

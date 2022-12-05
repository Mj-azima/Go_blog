package users

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type repoMock struct {
	GetResult Users
	GetError  error

	GetAllResult []Users
	GetAllError  error

	CreateResult uint
	CreateError  error

	UpdateError error

	DeleteResult Users
	DeleteError  error
}

func (r *repoMock) Get(id int) (Users, error) {
	return r.GetResult, r.GetError
}

func (r *repoMock) GetByEmail(email string) (Users, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repoMock) GetAll() ([]Users, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repoMock) Create(email string, password []byte) (uint, error) {
	return r.CreateResult, r.CreateError
}

func (r *repoMock) Edit(id int, email string, password []byte) error {
	//TODO implement me
	panic("implement me")
}

func (r *repoMock) Delete(id int) (Users, error) {
	//TODO implement me
	panic("implement me")
}

func TestServiceGet(t *testing.T) {
	id := uuid.New().ID()

	tests := map[string]struct {
		repo   Repo
		result Users
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				GetResult: Users{
					Model: gorm.Model{ID: uint(id)},
				},
				GetError: nil,
			},
			result: Users{
				Model: gorm.Model{ID: uint(id)},
			},
			err: nil,
		},
		"Not found from repo": {
			repo: &repoMock{
				GetResult: Users{},
				GetError:  ErrUserNotFound,
			},
			result: Users{},
			err:    ErrUserNotFound,
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
	userId := uuid.New().ID()
	//title := "some-title"

	//ac := UsersCreateUpdate{Title: title, Body: body}
	//ui := Users{
	//	Email:    "admin@example.com",
	//	Password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
	//}

	ur := Users{
		Model: gorm.Model{
			ID: uint(userId),
			//CreatedAt: time.Time{},
			//UpdatedAt: time.Time{},
		},
		Email:    "admin@example.com",
		Password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
	}

	tests := map[string]struct {
		repo     Repo
		result   Users
		email    string
		password []byte
		err      error
	}{
		"Happy path": {
			repo: &repoMock{
				CreateResult: uint(userId),
				CreateError:  nil,
				GetResult:    ur,
				GetError:     nil,
			},
			//input:  ac,
			email:    "admin@example.com",
			password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
			result:   ur,
			err:      nil,
		},
		"Not found from repo after create": {
			repo: &repoMock{
				CreateResult: uint(userId),
				CreateError:  nil,
				GetResult:    Users{},
				GetError:     ErrUserNotFound,
			},
			//input:  ac,
			email:    "admin@example.com",
			password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
			result:   Users{},
			err:      ErrUserNotFound,
		},
		"Creation failure": {
			repo: &repoMock{
				CreateResult: uint(userId),
				CreateError:  ErrUserCreate,
			},
			//input:  ac,
			email:    "admin@example.com",
			password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
			result:   Users{},
			err:      ErrUserCreate,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := New(test.repo)
			post, err := service.Create(test.email, test.password)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result, post)

		})
	}
}

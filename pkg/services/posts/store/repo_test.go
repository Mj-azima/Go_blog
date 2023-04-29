package store

import (
	"blog/pkg/services/posts"
	"blog/pkg/services/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

func ConnectToTempDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestPostRepoCreate2(t *testing.T) {

	body := "some-body"

	tests := map[string]struct {
		input  posts.Posts
		author users.Users
		body   string
		expect uint
		err    error
	}{
		"Happy path": {

			input:  posts.Posts{AuthorID: 1, Body: body},
			expect: 1,
			err:    nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {

			db, err := ConnectToTempDB()
			if err != nil {
				log.Fatal(err)
			}
			db.AutoMigrate(&posts.Posts{})
			repo := New(db)

			response, err := repo.Create(test.input.AuthorID, test.input.Body)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expect, response)

		})
	}
}

//Test Post Repository Get method
func TestPostRepoGet2(t *testing.T) {

	tests := map[string]struct {
		input  int
		expect posts.Posts
		err    error
	}{
		"Happy path": {

			input: 1,
			expect: posts.Posts{
				Model: gorm.Model{ID: 1},
			},
			err: nil,
		},
		"Unknown DB error": {

			input:  0,
			expect: posts.Posts{},
			err:    posts.ErrPostNotFound,
		},
		"Not found error": {

			input:  int(581684168),
			expect: posts.Posts{},
			err:    posts.ErrPostNotFound,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db, err := ConnectToTempDB()
			if err != nil {
				log.Fatal(err)
			}
			db.AutoMigrate(&posts.Posts{})
			repo := New(db)

			response, err := repo.Get(test.input)

			assert.Equal(t, test.err, err)
			//fmt.Println(test.err, err, "slgdnglsnglsdk")
			//fmt.Println(test.expect.ID, response.ID, "slgdnglsnglsdk")
			assert.Equal(t, test.expect.ID, response.ID)
		})
	}
}

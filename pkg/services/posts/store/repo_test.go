package store

import (
	"blog/pkg/services/posts"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

const (
	//selectPost      = `SELECT * FROM posts WHERE id=$1`
	//selectPost = "SELECT * FROM `posts` WHERE `posts`.`id` = ?"
	selectPost       = "SELECT * FROM `posts` WHERE `posts`.`id` = ? AND `posts`.`deleted_at` IS NULL ORDER BY `posts`.`id` LIMIT 1"
	selectPostEditor = "SELECT * FROM `post_editors` WHERE `post_editors`.`posts_id` = ?"
	selectManyPosts  = `SELECT * FROM posts LIMIT $1 OFFSET $2`
	//insertPost       = `INSERT INTO posts (author, body) VALUES ($1, $2) RETURNING id`
	//insertPost   = " INSERT INTO `posts` (created_at`,`updated_at`,`deleted_at`,`body`,`author_id`) VALUES (now(),now(),NULL,$1,$2) RETURNING `posts`.`id`"
	insertPost = `INSERT INTO "posts" ("created_at","updated_at","deleted_at","body",author_id) VALUES ($1,$2,$3,$4,$5) RETURNING "posts"."id"`

	insertAuthor = " INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`,`id`) VALUES ('now()','now()',NULL,$1,$2,$3)"
	//insertPost = " INSERT INTO `posts` (`created_at`,`updated_at`,`deleted_at`,`body`,`author_id`) VALUES (now(),now(),NULL,$1,$2)"
	updatePost = `UPDATE posts SET title = $1, body = $2, updated_at = now() WHERE id = $3`
)

//func TestPostRepoGet(t *testing.T) {
//	columns := []string{"ID"}
//	id := uuid.New().ID()
//	//now := time.Now()
//
//	mockResult := []driver.Value{int(id)}
//
//	tests := map[string]struct {
//		expectQueryArgs        []driver.Value
//		expectQueryResultRows  []*sqlmock.Rows
//		expectQueryResultError error
//		input                  int
//		expect                 posts.Posts
//		err                    error
//	}{
//		"Happy path": {
//			expectQueryArgs:        []driver.Value{id},
//			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
//			expectQueryResultError: nil,
//			input:                  int(id),
//			expect: posts.Posts{
//				Model: gorm.Model{ID: uint(id)},
//			},
//			err: nil,
//		},
//		"Unknown DB error": {
//			expectQueryArgs:        []driver.Value{id},
//			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
//			expectQueryResultError: errors.New("some-db-error"),
//			input:                  int(id),
//			expect:                 posts.Posts{},
//			err:                    posts.ErrPostNotFound,
//		},
//		"Not found error": {
//			expectQueryArgs:        []driver.Value{uint(581684168)},
//			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
//			expectQueryResultError: sql.ErrNoRows,
//			input:                  int(581684168),
//			expect:                 posts.Posts{},
//			err:                    posts.ErrPostNotFound,
//		},
//	}
//
//	for testName, test := range tests {
//		t.Run(testName, func(t *testing.T) {
//			db, mock, _ := sqlmock.New()
//			defer db.Close()
//
//			dialector := mysql.New(mysql.Config{
//				Conn:                      db,
//				DriverName:                "mysql",
//				SkipInitializeWithVersion: true,
//			})
//			gdb, err := gorm.Open(dialector, &gorm.Config{})
//
//			mock.ExpectQuery(regexp.QuoteMeta(selectPost)).WithArgs(test.expectQueryArgs...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)
//			mock.ExpectQuery(regexp.QuoteMeta(selectPostEditor)).WithArgs(test.expectQueryArgs...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)
//
//			repo := New(gdb)
//			response, err := repo.Get(test.input)
//
//			assert.Equal(t, test.err, err)
//			assert.Equal(t, test.expect.ID, response.ID)
//		})
//	}
//}

func TestPostRepoCreate(t *testing.T) {
	columns := []string{"id"}
	id := uuid.New().ID()
	//title := "some-title"
	body := "some-body"
	now := time.Now()
	//u := users.Users{
	//	Model: gorm.Model{
	//		ID: uint(id),
	//	},
	//
	//	Email:    "admin@example.com",
	//	Password: []byte("slgnsdlkgnsldkgmsdlkgsdlkg"),
	//}

	tests := map[string]struct {
		expectQueryArgs []driver.Value
		//expectQueryArgsUser    []driver.Value
		expectQueryResultRows  []*sqlmock.Rows
		expectQueryResultError error
		input                  posts.Posts
		//author                 users.Users
		//body                   string
		expect uint
		err    error
	}{
		"Happy path": {
			expectQueryArgs: []driver.Value{now, now, nil, body, int(id)},
			//expectQueryArgsUser:    []driver.Value{now, now, nil, u.Email, u.Password},
			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(int(id))},
			expectQueryResultError: nil,
			input:                  posts.Posts{AuthorID: int(id), Body: body},
			expect:                 uint(id),
			err:                    nil,
		},
		//"Create error": {
		//	expectQueryArgs:        []driver.Value{title, body},
		//	expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(id)},
		//	expectQueryResultError: errors.New("some-db-error"),
		//	input:                  posts.Posts{Author: u, Body: body},
		//	author:                 u,
		//	body:                   body,
		//	expect:                 0,
		//	err:                    posts.ErrPostCreate,
		//},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			dialector := mysql.New(mysql.Config{
				Conn:                      db,
				DriverName:                "mysql",
				SkipInitializeWithVersion: true,
			})
			gdb, err := gorm.Open(dialector, &gorm.Config{})
			mock.MatchExpectationsInOrder(false)
			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(insertPost)).WithArgs(test.expectQueryArgs...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)
			//mock.ExpectQuery(regexp.QuoteMeta(insertAuthor)).WithArgs(test.expectQueryArgsUser...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)
			mock.ExpectCommit()
			mock.ExpectRollback()

			repo := New(gdb)
			//fmt.Println(test.input.Author)
			//fmt.Println(test.input.Body)
			response, err := repo.Create(test.input.AuthorID, test.input.Body)

			fmt.Println(test.err)
			fmt.Println(err)
			fmt.Println(test.expect)
			fmt.Println(response)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expect, response)
		})
	}
}

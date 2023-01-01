package store

import (
	"blog/pkg/services/users"
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

const (
	selectPost = "SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
	//selectPostEditor = "SELECT * FROM `post_editors` WHERE `post_editors`.`posts_id` = ?"

)

func TestPostRepoGet(t *testing.T) {
	columns := []string{"ID"}
	id := uuid.New().ID()
	//now := time.Now()

	mockResult := []driver.Value{int(id)}

	tests := map[string]struct {
		expectQueryArgs        []driver.Value
		expectQueryResultRows  []*sqlmock.Rows
		expectQueryResultError error
		input                  int
		expect                 users.Users
		err                    error
	}{
		"Happy path": {
			expectQueryArgs:        []driver.Value{id},
			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
			expectQueryResultError: nil,
			input:                  int(id),
			expect: users.Users{
				Model: gorm.Model{ID: uint(id)},
			},
			err: nil,
		},
		"Unknown DB error": {
			expectQueryArgs:        []driver.Value{id},
			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
			expectQueryResultError: errors.New("some-db-error"),
			input:                  int(id),
			expect:                 users.Users{},
			err:                    users.ErrUserNotFound,
		},
		"Not found error": {
			expectQueryArgs:        []driver.Value{uint(581684168)},
			expectQueryResultRows:  []*sqlmock.Rows{sqlmock.NewRows(columns).AddRow(mockResult...)},
			expectQueryResultError: sql.ErrNoRows,
			input:                  int(581684168),
			expect:                 users.Users{},
			err:                    users.ErrUserNotFound,
		},
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

			mock.ExpectQuery(regexp.QuoteMeta(selectPost)).WithArgs(test.expectQueryArgs...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)
			//mock.ExpectQuery(regexp.QuoteMeta(selectPostEditor)).WithArgs(test.expectQueryArgs...).WillReturnError(test.expectQueryResultError).WillReturnRows(test.expectQueryResultRows...)

			repo := New(gdb)
			response, err := repo.Get(test.input)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expect.ID, response.ID)
		})
	}
}

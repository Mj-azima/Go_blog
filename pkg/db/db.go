package db

import (
	"blog/pkg/services/posts"
	"blog/pkg/services/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// GetConnection ...
func GetConnection(host string, port int, user, password, dbName string) (*gorm.DB, error) {
	if password == "" { // local DBs my not require a password
		password = `''`
	}
	//connStr := fmt.Sprintf("host=%s port=%d user=%s:password=%s@/dbname=%s sslmode=disable",
	//	host, port, user, password, dbName)
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	return nil, err
	//}

	//err = db.Ping()
	//if err != nil {
	//	return nil, err
	//}

	return db, nil
}

// Migrate ...
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&users.Users{}, &posts.Posts{})
	if err != nil {
		panic(err)
	}

	return nil
}

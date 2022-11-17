package database

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

var store *session.Store

func SetSession() {

	store = session.New(session.Config{
		Expiration: 120 * time.Second,
	})
	store.RegisterType(fiber.Map{})

}

func GetSession() *session.Store {
	return store
}

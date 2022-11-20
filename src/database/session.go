package database

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

type Session struct {
	store *session.Store
}

//var store *session.Store

func (s *Session) SetSession() {

	s.store = session.New(session.Config{
		Expiration: 120 * time.Second,
	})
	s.store.RegisterType(fiber.Map{})

}

func (s *Session) GetSession() *session.Store {
	return s.store
}

func (s *Session) Generate(c *fiber.Ctx, email string) error {
	store := s.store

	currSession, err := store.Get(c)
	if err != nil {
		return err
	}
	err = currSession.Regenerate()
	if err != nil {
		return err
	}
	currSession.Set("User", fiber.Map{"Email": email})
	err = currSession.Save()
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *Session) Get(c *fiber.Ctx) (fiber.Map, error) {
	store := s.store
	currSession, err := store.Get(c)
	if err != nil {
		return fiber.Map{}, err
	}
	user := currSession.Get("User")

	if user == nil {
		err := fmt.Errorf("do not found session")
		return fiber.Map{}, err
	}
	return user.(fiber.Map), nil
}

func (s *Session) Delete(c *fiber.Ctx) error {
	store := s.store
	currSession, err := store.Get(c)
	if err != nil {
		return err
	}
	user := currSession.Get("User")

	if user != nil {
		currSession.Delete("User")
	}
	return nil
}

package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"time"
)

var Instance *Session

func init() {
	Instance = new(Session)
}

type Session struct {
	store *session.Store
}

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

func (s *Session) Get(c *fiber.Ctx) (any, error) {
	store := s.store
	currSession, err := store.Get(c)
	if err != nil {
		return nil, err
	}
	user := currSession.Get("User")

	if user == nil {
		return nil, nil
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
		if err := currSession.Destroy(); err != nil {
			return err
		}
	}
	return nil
}

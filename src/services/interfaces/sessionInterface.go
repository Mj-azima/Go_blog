package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ISession interface {
	GetSession() *session.Store
	Generate(c *fiber.Ctx, email string) error
	Get(c *fiber.Ctx) (any, error)
	Delete(c *fiber.Ctx) error
}

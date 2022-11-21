package interfaces

import "github.com/gofiber/fiber/v2"

type Iauth interface {
	IsLogin(c *fiber.Ctx) (bool, error)
}

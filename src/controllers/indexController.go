package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//Index page controller
func Index(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{})
}

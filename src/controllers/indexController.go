package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//Index page controller
func Index(c *fiber.Ctx) error {
	isLogin, _ := IsLogin(c)
	return c.Render("index", fiber.Map{"isLogin": isLogin})
}

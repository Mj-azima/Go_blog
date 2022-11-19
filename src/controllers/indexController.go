package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//Index
type IndexStruct struct {
}

//Index page controller
func (i *IndexStruct) Index(c *fiber.Ctx) error {
	isLogin, _ := IsLogin(c)
	return c.Render("index", fiber.Map{"isLogin": isLogin})
}

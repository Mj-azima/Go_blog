package controllers

import (
	"blog/src/services"
	"github.com/gofiber/fiber/v2"
)

var auth *services.Authenticaion

func init() {
	auth = new(services.Authenticaion)
}

//Index
type IndexController struct {
}

//Index page controller
func (i *IndexController) Index(c *fiber.Ctx) error {

	isLogin, _ := auth.IsLogin(c)
	return c.Render("index", fiber.Map{"isLogin": isLogin})
}

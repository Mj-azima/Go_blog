package controllers

import "github.com/gofiber/fiber/v2"

//Create post Page controller
func CreatePostPage(c *fiber.Ctx) error {

	return c.Render("createPost", fiber.Map{})
}

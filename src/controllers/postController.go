package controllers

import (
	"blog/src/database"
	"blog/src/models"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
	"log"
)

//Create post Page controller
func CreatePostPage(c *fiber.Ctx) error {

	return c.Render("createPost", fiber.Map{})
}

//Create post request controller
func CreatePost(c *fiber.Ctx) error {
	payload := validators.Post{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	store := database.GetSession()

	currSession, err := store.Get(c)
	if err != nil {
		return err
	}
	usersess := currSession.Get("User").(fiber.Map)

	var user models.Users

	result := database.DBConn.Find(&user, "email = ?", usersess["Email"])

	if result.Error != nil {
		log.Fatal("not found a user")
		return result.Error
	}

	post := models.Posts{
		Auther: user.ID,
		Body:   payload.Body,
	}
	tx := database.DBConn.Create(&post)
	if tx.Error != nil {
		return tx.Error
	}

	return c.SendString("post created !")
}

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

//Update post page controller
func UpdatePostPage(c *fiber.Ctx) error {
	postId := c.Params("id")

	var post models.Posts
	result := database.DBConn.Find(&post, postId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return c.SendString("post not found")
	}

	return c.Render("updatePost", fiber.Map{"post": post})
}

//Update post request controller
func UpdatePost(c *fiber.Ctx) error {
	payload := validators.Post{}
	postId := c.Params("id")

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	var post models.Posts
	result := database.DBConn.Find(&post, postId)
	if result.Error != nil {
		return result.Error
	}
	result = database.DBConn.Model(&post).Update("Body", payload.Body)
	if result.Error != nil {
		return result.Error
	}

	return c.SendString("Post Updated!")
}

//Get all Posts page controller
func Posts(c *fiber.Ctx) error {
	var posts []models.Posts
	result := database.DBConn.Find(&posts)
	if result.Error != nil {
		return result.Error
	}

	return c.Render("postsList", fiber.Map{"posts": posts})
}

//Get signle post page controller
func SinglePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Posts

	if err := database.DBConn.First(&post, postId).Error; err != nil {
		return err
	}
	var user models.Users

	if err := database.DBConn.First(&user, post.Auther).Error; err != nil {
		return err
	}

	return c.Render("singlePost", fiber.Map{"post": post, "user": user})

}

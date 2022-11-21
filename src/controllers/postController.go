package controllers

import (
	"blog/src/database"
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type PostController struct {
}

//Create post Page controller
func (p *PostController) CreatePostPage(c *fiber.Ctx) error {

	return c.Render("createPost", fiber.Map{})
}

//Create post request controller
func (p *PostController) CreatePost(c *fiber.Ctx) error {
	payload := validators.Post{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	Session := services.Instance

	usersess, err := Session.Get(c) //Todo: session service
	if err != nil {
		return err
	}

	var user models.Users

	result := database.DBConn.Find(&user, "email = ?", (usersess).(fiber.Map)["Email"])

	if result.Error != nil {
		log.Fatal("not found a user")
		return result.Error
	}

	if err := postModel.Create(user.ID, payload.Body); err != nil {
		return err
	}

	return c.SendString("post created !")
}

//Update post page controller
func (p *PostController) UpdatePostPage(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := strconv.Atoi(postId) // type check
	post, result, err := postModel.Get(id)
	if err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return c.SendString("post not found")
	}

	return c.Render("updatePost", fiber.Map{"post": post})
}

//Update post request controller
func (p *PostController) UpdatePost(c *fiber.Ctx) error {
	payload := validators.Post{}
	postId := c.Params("id")

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	id, _ := strconv.Atoi(postId) // type check
	if err := postModel.Edit(id, payload.Body); err != nil {
		return err
	}

	return c.SendString("Post Updated!")
}

//Get all Posts page controller
func (p *PostController) Posts(c *fiber.Ctx) error {

	posts, err := postModel.GetAll()
	if err != nil {
		return err
	}
	return c.Render("postsList", fiber.Map{"posts": posts})
}

//Get signle post page controller
func (p *PostController) SinglePost(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := strconv.Atoi(postId) // type check
	post, _, err := postModel.Get(id)
	if err != nil {
		return err
	}
	var user models.Users

	if err := database.DBConn.First(&user, post.Auther).Error; err != nil {
		return err
	}

	return c.Render("singlePost", fiber.Map{"post": post, "user": user})

}

//Delete post request controller
func (p *PostController) DeletePost(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := strconv.Atoi(postId) // type check
	if err := postModel.Delete(id); err != nil {
		return err
	}

	return c.SendString("post Deleted!")
}

var postModel *repositories.Post

func init() {
	postModel = new(repositories.Post)
}

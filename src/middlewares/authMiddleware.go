package middlewares

import (
	"blog/src/repositories"
	"blog/src/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

//Check Logged in middleware
func RequireLogin(c *fiber.Ctx) error {
	//isLogin, err := controllers.IsLogin(c)
	isLogin, err := auth.IsLogin(c)
	if err != nil {
		log.Fatal(err)
	}

	if isLogin == false {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return c.Redirect("/login")
	}

	// If we got this far, the request is from a logged-in user.
	// Continue on to other middleware or routes.
	return c.Next()
}

//Check user is post author
func IsAuthor(c *fiber.Ctx) error {
	postId := c.Params("id")

	Session := services.Instance

	usersess, err := Session.Get(c)
	if err != nil {
		return err
	}

	//var user models.Users
	//database.DBConn.Find(&user, "email = ?", usersess.(fiber.Map)["Email"])

	user, _ := userModel.GetByEmail(usersess.(fiber.Map)["Email"].(string))

	//var post models.Posts
	//database.DBConn.Find(&post, "auther = ? AND id = ?", user.ID, postId)
	id, _ := strconv.Atoi(postId) // type check
	post, _ := postModel.GetByIdAndAuthor(user.ID, id)
	fmt.Println(post)

	return c.Next()
}

var auth *services.Authenticaion
var userModel *repositories.User
var postModel *repositories.Post

func init() {
	auth = new(services.Authenticaion)
	userModel = new(repositories.User)
	postModel = new(repositories.Post)
}

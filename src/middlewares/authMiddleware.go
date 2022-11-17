package middlewares

import (
	"blog/src/controllers"
	"blog/src/database"
	"blog/src/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

//Check Logged in middleware
func RequireLogin(c *fiber.Ctx) error {
	isLogin, err := controllers.IsLogin(c)
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

	store := database.GetSession()
	currSession, err := store.Get(c)
	if err != nil {
		return err
	}
	usersess := currSession.Get("User").(fiber.Map)

	var user models.Users
	database.DBConn.Find(&user, "email = ?", usersess["Email"])

	var post models.Posts
	database.DBConn.Find(&post, "auther = ? AND id = ?", user.ID, postId)

	return c.Next()
}

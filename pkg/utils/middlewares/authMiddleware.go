package middlewares

import (
	"blog/pkg/services/posts"
	"blog/pkg/services/users"
	"blog/pkg/utils/sessions"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type IAuthMiddleware interface {
	RequireLogin(c *fiber.Ctx) error
	IsAuthor(c *fiber.Ctx) error
}

type authMiddleware struct {
	userRepo    users.Repo
	postRepo    posts.Repo
	userService users.Service
}

func New(userRepo users.Repo, postRepo posts.Repo, userService users.Service) IAuthMiddleware {
	return &authMiddleware{userRepo, postRepo, userService}
}

//Check Logged in middlewares
func (a *authMiddleware) RequireLogin(c *fiber.Ctx) error {
	//isLogin, err := controllers.IsLogin(c)
	isLogin, err := a.userService.IsLogin(c)
	if err != nil {
		log.Fatal(err)
	}

	if isLogin == false {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return c.Redirect("/login")
	}

	// If we got this far, the request is from a logged-in user.
	// Continue on to other middlewares or routes.
	return c.Next()
}

//Check user is post author
func (a *authMiddleware) IsAuthor(c *fiber.Ctx) error {
	postId := c.Params("id")

	Session := sessions.Instance

	usersess, err := Session.Get(c)
	if err != nil {
		return err
	}

	usersession := usersess.(fiber.Map)
	email := usersession["Email"]
	userEmail, ok := email.(string)
	if !ok {
		return errors.New("type casting error (usersession)!")
	}

	user, _ := a.userRepo.GetByEmail(userEmail)

	id, _ := strconv.Atoi(postId) // type check
	_, _ = a.postRepo.GetByIdAndAuthor(user.ID, id)

	return c.Next()
}

package transport

import (
	"blog/pkg/services/posts"
	postStore "blog/pkg/services/posts/store"
	"blog/pkg/services/users"
	userStore "blog/pkg/services/users/store"
	"blog/pkg/utils/middlewares"
	"blog/pkg/utils/sessions"
	"blog/pkg/utils/validators"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type handler struct {
	PostService    posts.Service
	UserService    users.Service
	AuthMiddleware middlewares.IAuthMiddleware
}

func Activate(router *fiber.App, db *gorm.DB) {
	post := postStore.New(db)
	user := userStore.New(db)
	postService := posts.New(post)
	userService := users.New(user)
	authMiddleware := middlewares.New(user, post, userService)

	newHandler(router, postService, userService, authMiddleware)
}

func newHandler(router *fiber.App, ps posts.Service, us users.Service, am middlewares.IAuthMiddleware) {
	h := handler{
		PostService:    ps,
		UserService:    us,
		AuthMiddleware: am,
	}

	router.Get("/posts", h.posts)
	router.Get("/post", h.AuthMiddleware.RequireLogin, h.CreatePostPage)
	router.Post("/post", h.AuthMiddleware.RequireLogin, h.CreatePost)

	router.Put("/post/:id", h.AuthMiddleware.RequireLogin, h.AuthMiddleware.IsAuthor, h.Update)
	router.Get("/post/:id", h.AuthMiddleware.RequireLogin, h.AuthMiddleware.IsAuthor, h.updatePostPage)

	router.Get("/single-post/:id", h.singlePostPage)

	router.Post("/post/delete/:id", h.Delete)
}

func (h *handler) singlePostPage(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := strconv.Atoi(postId) // type check
	post, err := h.PostService.Get(id)
	if err != nil {
		return err
	}
	user, err := h.UserService.Get(post.AuthorID)

	return c.Render("posts/views/templates/singlePost", fiber.Map{"post": post, "user": user})
}

func (h *handler) posts(c *fiber.Ctx) error {
	allpost, err := h.PostService.GetAll()
	if err != nil {
		return err
	}

	return c.Render("posts/views/templates/postsList", fiber.Map{"posts": allpost})
}

func (h *handler) CreatePost(c *fiber.Ctx) error {
	payload := validators.Post{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	Session := sessions.Instance
	usersess, err := Session.Get(c) //Todo: session service
	if err != nil {
		return err
	}

	usersession := usersess.(fiber.Map)
	email := usersession["Email"].(string)

	user, err := h.UserService.GetByEmail(email)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		log.Fatal("not found a user")
	}

	err = h.PostService.Create(user, payload.Body)
	if err != nil {
		return err
	}

	return c.SendString("post created !")
}

func (h *handler) Update(c *fiber.Ctx) error {
	payload := validators.Post{}
	postId := c.Params("id")

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	id, _ := strconv.Atoi(postId) // type check
	err := h.PostService.Update(id, payload.Body)
	if err != nil {
		return err
	}

	return c.SendString("Post Updated!")
}

func (h *handler) Delete(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := strconv.Atoi(postId) // type check
	err := h.PostService.Delete(id)
	if err != nil {
		return err
	}
	return c.SendString("post Deleted!")
}

func (h *handler) CreatePostPage(c *fiber.Ctx) error {
	return c.Render("posts/views/templates/createPost", fiber.Map{})
}

func (h *handler) updatePostPage(c *fiber.Ctx) error {
	postId := c.Params("id")
	id, _ := strconv.Atoi(postId) // type check
	post, err := h.PostService.Get(id)
	if err != nil {
		return err
	}
	if post.ID == 0 {
		return c.SendString("post not found")
	}

	return c.Render("posts/views/templates/updatePost", fiber.Map{"post": post})
}

package transport

import (
	"blog/pkg/errors"
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
	"net/http"
	"strconv"
)

//post handler struct
type handler struct {
	PostService    posts.Service
	UserService    users.Service
	AuthMiddleware middlewares.IAuthMiddleware
}

//Activate function for use services
func Activate(router *fiber.App, db *gorm.DB) {
	post := postStore.New(db)
	user := userStore.New(db)
	postService := posts.New(post)
	userService := users.New(user)
	authMiddleware := middlewares.New(user, post, userService)

	newHandler(router, postService, userService, authMiddleware)
}

//new handler implement
func newHandler(router *fiber.App, ps posts.Service, us users.Service, am middlewares.IAuthMiddleware) {
	h := handler{
		PostService:    ps,
		UserService:    us,
		AuthMiddleware: am,
	}

	//Create post routers
	router.Get("/post", h.AuthMiddleware.RequireLogin, h.CreatePostPage)
	router.Post("/post", h.AuthMiddleware.RequireLogin, h.CreatePost)

	//Update & Delete post routers
	router.Post("/post/:id", h.AuthMiddleware.RequireLogin, h.Update)
	router.Get("/post/:id", h.AuthMiddleware.RequireLogin, h.updatePostPage)
	router.Post("/post/delete/:id", h.Delete)

	//Get single & list post routers
	router.Get("/single-post/:id", h.singlePostPage)
	router.Get("/posts", h.posts)

}

//Get single post page handler
func (h *handler) singlePostPage(c *fiber.Ctx) error {
	//Get param id
	postId := c.Params("id")

	//cast id & Get post
	id, _ := strconv.Atoi(postId) // type check
	post, err := h.PostService.Get(id)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{
			"status": status,
			"error":  appErr,
		})
	}

	//Get post's author
	user, err := h.UserService.Get(post.AuthorID)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{
			"status": status,
			"error":  appErr,
		})
	}

	return c.Render("posts/views/templates/singlePost", fiber.Map{"post": post, "user": user})
}

//Get all posts page handler
func (h *handler) posts(c *fiber.Ctx) error {
	//Get all posts
	allpost, err := h.PostService.GetAll()
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	return c.Render("posts/views/templates/postsList", fiber.Map{"posts": allpost})
}

//Get Crate post page handler
func (h *handler) CreatePost(c *fiber.Ctx) error {
	payload := validators.Post{}

	//Parse body context
	if err := c.BodyParser(&payload); err != nil {
		//return err
		return c.JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  errors.NewAppError(errors.BadRequest, errors.Descriptions[errors.BadRequest], ""),
		})
	}

	//Validate body context
	if err := validators.ValidateStruct(payload); err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	//Get user session
	Session := sessions.Instance
	usersess, err := Session.Get(c) //Todo: session service
	if err != nil {
		return err
	}

	//casting email from user session
	usersession := usersess.(fiber.Map)
	email := usersession["Email"].(string)

	//Get User By email
	user, err := h.UserService.GetByEmail(email)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}
	if user.ID == 0 {
		log.Fatal("not found a user")
	}

	//Create post with author & body
	_, err = h.PostService.Create(user, payload.Body)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	return c.SendString("post created !")
}

//Update Post handler
func (h *handler) Update(c *fiber.Ctx) error {
	payload := validators.Post{}

	//Get param id
	postId := c.Params("id")

	//Parse body context
	if err := c.BodyParser(&payload); err != nil {
		//return err
		return c.JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"error":  errors.NewAppError(errors.BadRequest, errors.Descriptions[errors.BadRequest], ""),
		})
	}

	//Validate body context
	if err := validators.ValidateStruct(payload); err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	//Get user session
	Session := sessions.Instance
	usersess, err := Session.Get(c)
	if err != nil {
		return err
	}

	//casting email from user session
	usersession := usersess.(fiber.Map)
	email := usersession["Email"].(string)

	//Get User By email
	user, err := h.UserService.GetByEmail(email)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	//Update post with author & body
	id, _ := strconv.Atoi(postId) // type check
	_, err = h.PostService.Update(id, payload.Body, user)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	return c.SendString("Post Updated!")
}

//Delete Post handler
func (h *handler) Delete(c *fiber.Ctx) error {
	//Get param id
	postId := c.Params("id")

	//Delete post by id
	id, _ := strconv.Atoi(postId) // type check
	_, err := h.PostService.Delete(id)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	return c.SendString("post Deleted!")
}

//Create post page handler
func (h *handler) CreatePostPage(c *fiber.Ctx) error {
	return c.Render("posts/views/templates/createPost", fiber.Map{})
}

//Update post page handler
func (h *handler) updatePostPage(c *fiber.Ctx) error {

	//Get param id
	postId := c.Params("id")

	//Get post by id
	id, _ := strconv.Atoi(postId) // type check
	post, err := h.PostService.Get(id)
	if err != nil {
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	//Check post exist
	if post.ID == 0 {
		return c.SendString("post not found")
	}

	return c.Render("posts/views/templates/updatePost", fiber.Map{"post": post})
}

// handleError allows us to map errors defined internally to appropriate HTTP error codes and JSON responses
func handleError(e error) (int, error) {
	switch e {
	case posts.ErrPostNotFound:
		return http.StatusNotFound, errors.NewAppError(errors.NotFound, e.Error(), "id")
	case posts.ErrPostUpdate:
		fallthrough
	case posts.ErrPostCreate:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, "unable to create/update post", "")
	default:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, e.Error(), "unknown")
	}
}

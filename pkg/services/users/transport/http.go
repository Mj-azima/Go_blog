package transport

import (
	"blog/pkg/services/users"
	"blog/pkg/services/users/store"
	"blog/pkg/utils/cryptography"
	"blog/pkg/utils/sessions"
	"blog/pkg/utils/validators"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	UserService users.Service
}

func Activate(router *fiber.App, db *gorm.DB) {
	userService := users.New(store.New(db))

	newHandler(router, userService)
}

func newHandler(router *fiber.App, us users.Service) {
	h := handler{
		UserService: us,
	}

	router.Post("/login", h.login)
	router.Get("/login", h.loginPage)
	router.Get("/register", h.registerPage)
	router.Post("/register", h.register)
	router.Post("/logout", h.logout)

}

func (h *handler) login(c *fiber.Ctx) error {

	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	user, err := h.UserService.GetByEmail(payload.Email)
	if err != nil {
		return err
	}

	if err := cryptography.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	Session := sessions.Instance

	if err := Session.Generate(c, user.Email); err != nil {
		return err
	} //Todo: session service

	return c.Redirect("/")
}

func (h *handler) register(c *fiber.Ctx) error {
	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	passwd, err := cryptography.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return err
	}

	if err := h.UserService.Create(payload.Email, passwd); err != nil {
		return err
	}
	return c.Redirect("/login")

}

func (h *handler) loginPage(c *fiber.Ctx) error {
	return c.Render("users/views/templates/login", fiber.Map{})
}

func (h *handler) registerPage(c *fiber.Ctx) error {
	return c.Render("users/views/templates/register", fiber.Map{})
}

func (h *handler) logout(c *fiber.Ctx) error {
	Session := sessions.Instance

	if err := Session.Delete(c); err != nil {
		return err
	} //Todo: session service

	return c.Redirect("/")
}

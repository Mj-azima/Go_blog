package transport

import (
	"blog/pkg/errors"
	"blog/pkg/services/users"
	"blog/pkg/services/users/store"
	"blog/pkg/utils/cryptography"
	"blog/pkg/utils/sessions"
	"blog/pkg/utils/validators"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
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
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	}

	if err := cryptography.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		//return c.JSON(fiber.Map{
		//	"message": "incorrect password",
		//})
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
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
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
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
		//return err
		status, appErr := handleError(err)
		return c.JSON(fiber.Map{"status": status, "err": appErr})
	} //Todo: session service

	return c.Redirect("/")
}

func handleError(e error) (int, error) {
	switch e {
	case users.ErrUserNotFound:
		return http.StatusNotFound, errors.NewAppError(errors.NotFound, e.Error(), "id")
	case users.ErrUserUpdate:
		fallthrough
	case users.ErrUserCreate:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, "unable to create/update post", "")
	case cryptography.ErrIncorrectPasswordError:
		return http.StatusForbidden, errors.NewAppError(errors.BadRequest, "incorrect password Error", "")
	default:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, e.Error(), "unknown")
	}
}

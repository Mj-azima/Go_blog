package controllers

import (
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
)

//Authentication
type AuthenticaionController struct {
}

//Register page controller
func (a *AuthenticaionController) RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

//Register request controller
func (a *AuthenticaionController) Register(c *fiber.Ctx) error {
	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	bcrypt := new(services.Bcrypt)
	passwd, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return err
	}

	if err := userModel.Create(payload.Email, passwd); err != nil {
		return err
	}

	return c.Redirect("/login")
}

//Login page controller
func (a *AuthenticaionController) LoginPage(c *fiber.Ctx) error {

	return c.Render("login", fiber.Map{})
}

//Login request controller
func (a *AuthenticaionController) Login(c *fiber.Ctx) error {

	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	user, err := userModel.GetByEmail(payload.Email)
	if err != nil {
		return err
	}

	bcrypt := new(services.Bcrypt)
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	Session := services.Instance

	if err := Session.Generate(c, user.Email); err != nil {
		return err
	} //Todo: session service

	return c.Redirect("/")

}

//Logout request controller
func (a *AuthenticaionController) Logout(c *fiber.Ctx) error {

	Session := services.Instance

	if err := Session.Delete(c); err != nil {
		return err
	} //Todo: session service

	return c.Redirect("/")
}

var userModel *repositories.User

func init() {
	userModel = new(repositories.User)
}

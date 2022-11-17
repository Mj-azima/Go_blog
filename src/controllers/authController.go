package controllers

import (
	"blog/src/database"
	"blog/src/models"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func Register(c *fiber.Ctx) error {
	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	passwd, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	user := models.Users{
		Email:    payload.Email,
		Password: passwd,
	}
	tx := database.DBConn.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return c.Redirect("/login")
}

func LoginPage(c *fiber.Ctx) error {

	return c.Render("login", fiber.Map{})
}

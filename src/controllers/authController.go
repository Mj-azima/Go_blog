package controllers

import (
	"blog/src/database"
	"blog/src/models"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func Login(c *fiber.Ctx) error {

	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	//var dbUser models.Users
	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}
	var user models.Users

	result := database.DBConn.Find(&user, "email = ?", payload.Email)

	if result.Error != nil {
		log.Fatal("not found a user")
		return result.Error
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	store := database.GetSession()

	currSession, err := store.Get(c)
	if err != nil {
		return err
	}
	err = currSession.Regenerate()
	if err != nil {
		return err
	}
	currSession.Set("User", fiber.Map{"Email": user.Email})
	err = currSession.Save()
	if err != nil {
		panic(err)
	}

	return c.Redirect("/")

}

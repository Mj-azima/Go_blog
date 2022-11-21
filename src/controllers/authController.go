package controllers

import (
	"blog/src/database"
	"blog/src/repositories"
	"blog/src/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//Authentication
type AuthenticationStruct struct {
}

//Register page controller
func (a *AuthenticationStruct) RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

//Register request controller
func (a *AuthenticationStruct) Register(c *fiber.Ctx) error {
	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}

	passwd, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	//user := models.Users{
	//	Email:    payload.Email,
	//	Password: passwd,
	//}
	//tx := database.DBConn.Create(&user)
	//if tx.Error != nil {
	//	return tx.Error
	//}

	if err := userModel.Create(payload.Email, passwd); err != nil {
		return err
	}

	return c.Redirect("/login")
}

//Login page controller
func (a *AuthenticationStruct) LoginPage(c *fiber.Ctx) error {

	return c.Render("login", fiber.Map{})
}

//Login request controller
func (a *AuthenticationStruct) Login(c *fiber.Ctx) error {

	payload := validators.User{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	//var dbUser models.Users
	if err := validators.ValidateStruct(payload); err != nil {
		return err
	}
	//var user models.Users
	//
	//result := database.DBConn.Find(&user, "email = ?", payload.Email)
	//
	//if result.Error != nil {
	//	log.Fatal("not found a user")
	//	return result.Error
	//}

	user, err := userModel.GetByEmail(payload.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(payload.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//store := database.GetSession()
	//
	//currSession, err := store.Get(c)
	//if err != nil {
	//	return err
	//}
	//err = currSession.Regenerate()
	//if err != nil {
	//	return err
	//}
	//currSession.Set("User", fiber.Map{"Email": user.Email})
	//err = currSession.Save()
	//if err != nil {
	//	panic(err)
	//}

	Session := database.Instance

	if err := Session.Generate(c, user.Email); err != nil {
		return err
	}

	return c.Redirect("/")

}

//Logout request controller
func (a *AuthenticationStruct) Logout(c *fiber.Ctx) error {
	//store := database.GetSession()
	//currSession, err := store.Get(c)
	//if err != nil {
	//	return err
	//}
	//user := currSession.Get("User")

	//user,err := Session.Get(c)
	//if err != nil {
	//	return err
	//}

	//if user != nil {
	//	currSession.Delete("User")
	//}

	//err = currSession.Save()
	//if err != nil {
	//	panic(err)
	//}
	Session := database.Instance

	if err := Session.Delete(c); err != nil {
		return err
	}

	return c.Redirect("/")
}

//IsLogin service
func IsLogin(c *fiber.Ctx) (bool, error) {
	//store := database.GetSession()
	//currSession, err := store.Get(c)
	//if err != nil {
	//	return false, err
	//}
	//user := currSession.Get("User")
	Session := database.Instance

	user, err := Session.Get(c)
	if err != nil {
		return false, err
	}

	if user == nil {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return false, nil
	}
	return true, nil
}

//var Session *database.Session
//
//func init() {
//	Session = new(database.Session)
//	Session.SetSession()
//}

var userModel *repositories.User

func init() {
	userModel = new(repositories.User)
}

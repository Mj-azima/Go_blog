package services

import "github.com/gofiber/fiber/v2"

type Authenticaion struct {
}

//IsLogin service
func (a *Authenticaion) IsLogin(c *fiber.Ctx) (bool, error) {
	Session := Instance

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

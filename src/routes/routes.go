package routes

import (
	"blog/src/controllers"
	"github.com/gofiber/fiber/v2"
)

//setRoute config
func SetUpRoutes(app *fiber.App) {
	//index page route
	app.Get("/", controllers.Index)

}

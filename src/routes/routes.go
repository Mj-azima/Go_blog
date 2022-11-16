package routes

import (
	"blog/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", controllers.Index)

}

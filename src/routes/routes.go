package routes

import (
	"blog/src/controllers"
	"blog/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

//SetRoute config
func SetUpRoutes(app *fiber.App) {
	//Index page route
	app.Get("/", controllers.Index)

	//Register page route
	app.Get("/register", controllers.RegisterPage)
	//Register request route
	app.Post("/register", controllers.Register)

	//Login page route
	app.Get("/login", controllers.LoginPage)
	//Login request route
	app.Post("/login", controllers.Login)
	//Logout request route
	app.Post("/logout", controllers.Logout)

	//Create Post page route
	app.Get("/post", middlewares.RequireLogin, controllers.CreatePostPage)
	//Create Post request route
	app.Post("/post", middlewares.RequireLogin, controllers.CreatePost)

	//Update Post page route
	app.Get("/post/:id", middlewares.RequireLogin, middlewares.IsAuthor, controllers.UpdatePostPage)
	//Update Post request route
	app.Post("/post/:id", middlewares.RequireLogin, middlewares.IsAuthor, controllers.UpdatePost)

	//Get all posts page
	app.Get("/posts", controllers.Posts)
	//Get single post page
	app.Get("/singlePost/:id", controllers.SinglePost)
}

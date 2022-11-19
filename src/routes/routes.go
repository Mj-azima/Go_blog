package routes

import (
	"blog/src/controllers"
	"blog/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

//SetRoute config
func SetUpRoutes(app *fiber.App) {

	//Index
	indStr := new(controllers.IndexStruct)
	//Index page route
	app.Get("/", indStr.Index)

	//Authentication
	authStruct := new(controllers.AuthenticationStruct)
	//Register page route
	app.Get("/register", authStruct.RegisterPage)
	//Register request route
	app.Post("/register", authStruct.Register)

	//Login page route
	app.Get("/login", authStruct.LoginPage)
	//Login request route
	app.Post("/login", authStruct.Login)
	//Logout request route
	app.Post("/logout", authStruct.Logout)

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

	//Delete post request
	app.Post("/post/delete/:id", middlewares.RequireLogin, middlewares.IsAuthor, controllers.DeletePost)
}

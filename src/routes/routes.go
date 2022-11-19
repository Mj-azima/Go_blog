package routes

import (
	"blog/src/controllers"
	"blog/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

//SetRoute config
func SetUpRoutes(app *fiber.App) {

	//Index
	indexStruct := new(controllers.IndexStruct)
	//Index page route
	app.Get("/", indexStruct.Index)

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

	//Post
	postStruct := new(controllers.PostStruct)
	//Create Post page route
	app.Get("/post", middlewares.RequireLogin, postStruct.CreatePostPage)
	//Create Post request route
	app.Post("/post", middlewares.RequireLogin, postStruct.CreatePost)

	//Update Post page route
	app.Get("/post/:id", middlewares.RequireLogin, middlewares.IsAuthor, postStruct.UpdatePostPage)
	//Update Post request route
	app.Post("/post/:id", middlewares.RequireLogin, middlewares.IsAuthor, postStruct.UpdatePost)

	//Get all posts page
	app.Get("/posts", postStruct.Posts)
	//Get single post page
	app.Get("/singlePost/:id", postStruct.SinglePost)

	//Delete post request
	app.Post("/post/delete/:id", middlewares.RequireLogin, middlewares.IsAuthor, postStruct.DeletePost)
}

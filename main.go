package main

import (
	"blog/src/database"
	"blog/src/routes"
	"blog/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	templateEngineDir := os.Getenv("TEMPLATE_ENGINE_DIR")
	engine := html.New(templateEngineDir, ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.ConnectDb()
	Session := services.Instance
	Session.SetSession()

	routes.SetUpRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

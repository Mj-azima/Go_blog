package main

import (
	"blog/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{})

	database.ConnectDb()

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

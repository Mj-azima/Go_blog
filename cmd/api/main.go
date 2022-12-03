package main

import (
	"blog/pkg/api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {

	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	runMigration, err := strconv.ParseBool(os.Getenv("RUN_MIGRATION"))
	if err != nil {
		panic(err)
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	api.Start(&api.Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       dbPort,
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		RunMigration: runMigration,

		AppHost: os.Getenv("HOST"),
		AppPort: port,
	})
}

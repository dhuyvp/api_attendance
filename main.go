package main

import (
	"attendance_user/pkg/routes"
	"attendance_user/platform/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to loading .env file")
	}

	db, err := database.ConnectionDB()
	if err != nil {
		log.Fatal(fmt.Println(err))
	} else {
		fmt.Println("Database connected!")
	}

	app := fiber.New()

	routes.PublicRoutes(app, db)

	if err := app.Listen(":1206"); err != nil {
		log.Printf("Oops ... Server is not running! Reason: %v", err)
	}

}

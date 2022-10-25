package routes

import (
	"attendance_user/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicRoutes(app *fiber.App, db *sqlx.DB) {
	route := app.Group("/api")

	var idSelect int = 1
	var idInsert int = 1

	route.Post("/insert", controllers.InsertUser(db, idInsert))
	route.Get("/select", controllers.SelectUser(db, idSelect))
	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}

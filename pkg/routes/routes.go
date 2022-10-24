package routes

import (
	"attendance_user/app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicRoutes(app *fiber.App, db *sqlx.DB) {
	route := app.Group("/api")

	route.Post("/insert", controllers.InsertUser(db))
	route.Get("/select", controllers.SelectUser(db, 1))
	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}

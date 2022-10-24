package controllers

import (
	"attendance_user/app/models"
	"attendance_user/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SelectUser(db *sqlx.DB, userid int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryDb := "SELECT * FROM User WHERE UserID = $1 ORDER BY Id DESC LIMIT 1"

		var queryResult []models.User
		err := db.Select(&queryResult, queryDb, userid)

		if err != nil {
			log.Println("Error to select database", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       queryResult,
		})
	}
}

func InsertUser(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// UserID want to insert
		var userid = 1
		ReceiveCoin := []int{1, 1, 1, 1, 1, 1, 10}

		// Select to get last id with UserID = userid
		querySelect := "SELECT * FROM Attendance WHERE UserID = ? ORDER BY Id DESC LIMIT 1"

		var result []models.User
		err := db.Select(&result, querySelect, userid)
		if err != nil {
			log.Println("Error to select database before insert data!", err)
			return c.Status(fiber.StatusNotAcceptable).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusNotAcceptable,
				//Data:       user,
			})
		}

		// Information of user that i insert to database
		user := models.User{
			UserID:   userid,
			CountDay: 1,
			Coin:     1,
		}

		if result != nil {
			user.Coin = result[0].Coin + 1

			LastDate := result[0].CreatedAt

			if result != nil && utils.CheckAttendance(LastDate) {
				user.CountDay = (result[0].CountDay + 1) % 8
				if user.CountDay == 0 {
					user.CountDay = 1
				}
				user.Coin += (ReceiveCoin[user.CountDay-1] - 1)
			}
		}

		// Insert data to database
		tableName := "Attendance"
		queryDb := "INSERT INTO " + tableName + " (UserID, CountDay, Coin) VALUES (?, ?, ?)"

		_, errInsert := db.Exec(queryDb, user.UserID, user.CountDay, user.Coin)
		if errInsert != nil {
			log.Println("Error to select database", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       user,
		})
	}
}

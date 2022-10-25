package controllers

import (
	"attendance_user/app/models"
	"attendance_user/pkg/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SaveBeforeResponse(user models.User) []models.User {
	ReceiveCoin := []int{1, 2, 3, 4, 5, 6, 10}

	Result := [7]models.User{}

	val := utils.CheckAttendance(user.CreatedAt)

	for i := 0; i < 7; i++ {
		Result[i].IsAttendance = false
		Result[i].Coin = ReceiveCoin[i]
		Result[i].Day = "Ngày " + strconv.Itoa(i+1)
	}

	if val == 0 {
		for i := 0; i < user.CountDay; i++ {
			Result[i].IsAttendance = true
		}
		Result[user.CountDay-1].Day = "Hôm nay"
		return Result[:]
	}
	if val == 1 {
		for i := 0; i < user.CountDay-1; i++ {
			Result[i].IsAttendance = true
		}
		Result[user.CountDay-1].Day = "Hôm nay"
		return Result[:]
	}

	Result[0].Day = "Hôm nay"
	return Result[:]
}

func SelectUser(db *sqlx.DB, userid int) fiber.Handler {
	return func(c *fiber.Ctx) error {

		queryDb := "SELECT * FROM Attendance WHERE UserID = ? ORDER BY Id DESC LIMIT 1"

		queryResult := []models.User{}
		err := db.Select(&queryResult, queryDb, userid)

		if err != nil {
			log.Println("Error to select database", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		Result := SaveBeforeResponse(queryResult[0])

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       Result,
		})
	}
}

func InsertUser(db *sqlx.DB, userid int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ReceiveCoin := []int{1, 2, 3, 4, 5, 6, 10}

		// Select to get last id with UserID = userid
		querySelect := "SELECT * FROM Attendance WHERE UserID = ? ORDER BY Id DESC LIMIT 1"

		var result []models.User
		err := db.Select(&result, querySelect, userid)
		if err != nil {
			log.Println("Error to select database before insert data!", err)
			return c.Status(fiber.StatusNotAcceptable).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusNotAcceptable,
			})
		}

		// Information of user that i insert to database
		user := models.User{
			UserID:   userid,
			CountDay: 1,
			Coin:     ReceiveCoin[0],
		}

		if result != nil {
			user.Coin = result[0].Coin + ReceiveCoin[0]
			LastDate := result[0].CreatedAt

			var val int = utils.CheckAttendance(LastDate)

			// Attendance on the same day
			if val == 0 {
				user.CountDay = result[0].CountDay
				user.Coin = result[0].Coin
			}

			// Attendance yesterday
			if val == 1 {
				user.CountDay = (result[0].CountDay + 1) % 8

				if user.CountDay == 0 {
					user.CountDay = 1
				}

				user.Coin += (ReceiveCoin[user.CountDay-1] - ReceiveCoin[0])
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

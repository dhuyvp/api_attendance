package utils

import (
	"time"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"err_message"`
	Data       interface{} `json:"data"`
}

func CheckAttendance(LastDate string) int {
	timeInsert := time.Now()
	InsertFormat := timeInsert.Format("2006-01-02 15:04:05")

	tmp := timeInsert.AddDate(0, 0, -1)
	str := tmp.Format("2006-01-02 15:04:05")

	valLastDay := LastDate[:10]    // LastInsertDate
	newValDay := str[:10]          // Date after add -1 day
	valInsert := InsertFormat[:10] // Date now

	if valInsert == valLastDay {
		return 0
	}

	if newValDay == valLastDay {
		return 1
	}

	return 2
}

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

func ConvertToInt(x byte, y byte) int {
	result := int(x-'0')*10 + int(y-'0')
	return result
}

func CheckAttendance(LastDate string) bool {
	timeInsert := time.Now()
	str := timeInsert.Format("2006-01-02 15:04:05")

	val := ConvertToInt(LastDate[8], LastDate[9])
	newVal := ConvertToInt(str[8], str[9])

	return (newVal-val >= 0 && newVal-val < 2)
}

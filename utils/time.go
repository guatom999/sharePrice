package utils

import "time"

func LoadLocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func GetLocalDate() time.Weekday {
	bkkTime, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(bkkTime)

	today := now.Weekday()
	return today
}

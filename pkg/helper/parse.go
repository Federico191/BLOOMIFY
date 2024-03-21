package helper

import (
	"time"
)

func ParseDateTime(dayStr, timeStr string) (time.Time, error) {
	dateFormat := "02-01-2006"

	timeFormat := "15.04"

	dateTimeStr := dayStr + " " + timeStr

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return time.Time{}, err
	}

	dateTime, err := time.ParseInLocation(dateFormat+" "+timeFormat, dateTimeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

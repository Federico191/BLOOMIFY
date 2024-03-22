package helper

import (
	"time"
)

func ParseDateTime(dayStr, timeStr string) (time.Time, error) {
	dateFormat := "2006-01-02"

	timeFormat := "15.04"

	dateTimeStr := dayStr + " " + timeStr

	loc, err := time.LoadLocation("Asia/Jakarta") // Atur sesuai dengan zona waktu yang diinginkan
	if err != nil {
		return time.Time{}, err
	}

	dateTime, err := time.ParseInLocation(dateFormat+" "+timeFormat, dateTimeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

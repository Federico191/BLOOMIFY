package helper

import (
	"time"
)

func ParseDateTime(dayStr, timeStr string) (time.Time, error) {
	// Format string untuk tanggal
	dateFormat := "2006-01-02" // Format: yyyy-mm-dd

	// Format string untuk waktu
	timeFormat := "15.04" // Format: HH:mm (24 jam)

	// Gabungkan string tanggal dan waktu
	dateTimeStr := dayStr + " " + timeStr

	// Set zona waktu yang diharapkan
	loc, err := time.LoadLocation("Asia/Jakarta") // Atur sesuai dengan zona waktu yang diinginkan
	if err != nil {
		return time.Time{}, err
	}

	// Parsing string menjadi time.Time dengan zona waktu yang diharapkan
	dateTime, err := time.ParseInLocation(dateFormat+" "+timeFormat, dateTimeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

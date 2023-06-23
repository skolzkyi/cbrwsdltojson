package helpers

import (
	"strings"
	"time"
)

func OnlyDate(inputDateTime time.Time) time.Time {
	year := inputDateTime.Year()
	month := inputDateTime.Month()
	day := inputDateTime.Day()
	location := inputDateTime.Location()

	OnlyDate := time.Date(year, month, day, 0, 0, 0, 0, location)

	return OnlyDate
}

func DateStartTime(inputDateTime time.Time) time.Time {
	return OnlyDate(inputDateTime)
}

func DateEndTime(inputDateTime time.Time) time.Time {
	year := inputDateTime.Year()
	month := inputDateTime.Month()
	day := inputDateTime.Day()
	location := inputDateTime.Location()

	DateEndTime := time.Date(year, month, day, 23, 59, 59, 0, location)

	return DateEndTime
}

func DateBetweenInclude(inputDateTime time.Time, startTime time.Time, endTime time.Time) bool {
	aft := inputDateTime.After(startTime) || inputDateTime.Equal(startTime)
	bf := inputDateTime.Before(endTime) || inputDateTime.Equal(endTime)
	return aft && bf
}

func StringBuild(input ...string) string {
	var sb strings.Builder

	for _, str := range input {
		sb.WriteString(str)
	}

	return sb.String()
}

package utils

import "time"

func IsEvenWeek() bool {
	_, week := time.Now().ISOWeek()
	return week%2 == 0
}

func GetDayOfWeek() int {
	return int(time.Now().Weekday())
}

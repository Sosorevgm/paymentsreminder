package internal

import (
	"time"
)

func CheckTime(
	currentWeekday time.Weekday,
	currentHour int,
	currentMinute int,
	requiredWeekdays []time.Weekday,
	requiredHour int,
	requiredMinute int,
) bool {
	if containsWeekday(currentWeekday, requiredWeekdays) {
		if currentHour == requiredHour {
			if currentMinute == requiredMinute {
				return true
			}
		}
	}
	return false
}

func containsWeekday(currentWeekday time.Weekday, requiredWeekdays []time.Weekday) bool {
	result := false
	for _, iterableDay := range requiredWeekdays {
		if iterableDay == currentWeekday {
			result = true
		}
	}
	return result
}

package util

import "time"

const layout = "2006-01-02T15:04:05"

// StringToTime function to convert a string in a time object
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)
	return convertedTime
}

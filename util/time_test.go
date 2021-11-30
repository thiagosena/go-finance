package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:11:12")

	if convertedTime.Year() != 2019 {
		t.Errorf("The expected year was 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("The expected month was 02")
	}

	if convertedTime.Day() != 12 {
		t.Errorf("The expected day was 12")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("The expected day was 10")
	}

	if convertedTime.Minute() != 11 {
		t.Errorf("The expected day was 11")
	}

	if convertedTime.Second() != 12 {
		t.Errorf("The expected day was 12")
	}
}

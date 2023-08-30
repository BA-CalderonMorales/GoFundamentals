package utils

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Brandon"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

// Date represents a date
type Date struct {
	day, month, year int // properties not exported
}

// SetDay sets the day of the Date.
//
// Parameters:
// - day: an integer representing the day value.
//
// Returns:
// - error: an error if the day value is incorrect.
func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("Incorrect day value")
	}
	d.day = day
	return nil
}

// SetMonth sets the month of the Date.
//
// month: an integer representing the month (1-12).
// error: returns an error if the month value is incorrect.
func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("Incorrect month value")
	}
	d.month = month
	return nil
}

// SetYear sets the year of the Date.
//
// year: the year value to set.
// error: returns an error if the year value is incorrect.
func (d *Date) SetYear(year int) error {
	if (year < 1875) || (year > time.Now().Year()) {
		return errors.New("Incorrect year value")
	}
	d.year = year
	return nil
}

// Day returns the day of the month for the given Date.
//
// No parameters.
// Returns an integer representing the day of the month.
func (d *Date) Day() int {
	return d.day
}

// Month returns the month of the Date.
//
// No parameters.
// Returns an integer representing the month.
func (d *Date) Month() int {
	return d.month
}

// Year returns the year of the Date.
//
// No parameters.
// Returns an integer.
func (d *Date) Year() int {
	return d.year
}

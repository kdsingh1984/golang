// Package calendar to set date
package calendar

import (
	"errors"
)

// Date structure
type Date struct {
	day   int
	month int
	year  int
}

// SetYear set year in date structure
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth set month in date structure
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 13 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay set day in date structure
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

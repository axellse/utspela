// Handles the asttime format used by utspela
// https://github.com/pyrretsoftware/asttime
package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	FormatExample = "Sun 31/8/2025 21:30:16"
)

var WeekDays = []string{
	"mon", "tue", "wed", "thu", "fri", "sat", "sun",
}

var WeekDaysTable = map[time.Weekday]string{
	time.Monday:    "mon",
	time.Tuesday:   "tue",
	time.Wednesday: "wed",
	time.Thursday:  "thu",
	time.Friday:    "fri",
	time.Saturday:  "sat",
	time.Sunday:    "sun",
}

type AstTime struct {
	weekDay string
	dayOfMonth int
	month int
	year int
	hour int
	minute int
	second int
}

//ParseString parses a string in the asttime format, returning any errors it encounters
func ParseString(f string) (AstTime, error) {
	sections := strings.Split(f, " ")
	at := AstTime{
		dayOfMonth: -1,
		month: -1,
		year: -1,
		hour: -1,
		minute: -1,
		second: -1,
	}

	for _, s  := range sections {
		if len(s) == 3 {
			if !slices.Contains(WeekDays, strings.ToLower(s)) {
				return AstTime{}, errors.New("invalid weekday specified: " + s)
			}

			at.weekDay = s
		} else if strings.Contains(s, "/") {
			if strings.Count(s, "/") != 2 {
				return AstTime{}, errors.New("date component needs to have two slashes")
			}
			c := strings.Split(s, "/")
			day, month, year := c[0], c[1], c[2]
			dayint, monthint, yearint := -1, -1, -1

			if day != "*" {
				_da, err := strconv.Atoi(day)
				dayint = _da
				if err != nil {return AstTime{}, errors.New("invalid number on date component, day")}
			} else {dayint = -1}

			if month != "*" {
				_ma, err := strconv.Atoi(month)
				monthint = _ma
				if err != nil {return AstTime{}, errors.New("invalid number on date component, month")}
			} else {monthint = -1}
			
			if year != "*" {
				_ya, err := strconv.Atoi(year)
				yearint = _ya
				if err != nil {return AstTime{}, errors.New("invalid number on date component, year")}
			} else {yearint = -1}

			if dayint > 31 || monthint > 12 {
				return AstTime{}, errors.New("day or month over max value, day max is 31, month max is 12")
			}

			at.dayOfMonth, at.month, at.year = dayint, monthint, yearint
		} else if strings.Contains(s, ":") {
			if strings.Count(s, ":") != 2 {
				return AstTime{}, errors.New("time component needs to have two colons")
			}

			c := strings.Split(s, ":")
			hour, min, sec := c[0], c[1], c[2]
			hourint, minint, secint := -1, -1, -1

			if hour != "*" {
				_ha, err := strconv.Atoi(hour)
				hourint = _ha
				if err != nil {return AstTime{}, errors.New("invalid number on time component, hour")}
			} else {hourint = -1}

			if min != "*" {
				_ma, err := strconv.Atoi(min)
				minint = _ma
				if err != nil {return AstTime{}, errors.New("invalid number on time component, minute")}
			} else {minint = -1}
			
			if sec != "*" {
				_si, err := strconv.Atoi(sec)
				secint = _si
				if err != nil {return AstTime{}, errors.New("invalid number on time component, second")}
			} else {secint = -1}

			if hourint > 23 || minint > 59 || secint > 59 {
				return AstTime{}, errors.New("hour, min or sec over max value, hour max is 23, min and sec max is 59")
			}

			at.hour, at.minute, at.second = hourint, minint, secint
		}
	}

	return at, nil
}

//ongoing returns whether or not the current time matches the asttime
func (a AstTime) Ongoing() bool {
	hour, min, sec := time.Now().Clock()
	year, month, day := time.Now().Date()
	weekday := time.Now().Weekday()

	if (hour == a.hour || a.hour == -1) &&
	 (min == a.minute || a.minute == -1) &&
	 (sec == a.second || a.second == -1) &&
	 (year == a.year || a.year == -1) &&
	 (month == time.Month(a.month) || a.month == -1) &&
	 (day == a.dayOfMonth || a.dayOfMonth == -1) &&
	 (WeekDaysTable[weekday] == strings.ToLower(a.weekDay) || a.weekDay == "") {
		return true
	}
	return false
}
package utils

import (
	"fmt"
	"strconv"
)

// IsValidYear checks if given string is a 4 digits number, like "1234" (not "-123", "123", or "12345").
func IsValidYear(num int) bool {
	return num > 999 && num < 10_000
}

// SliceYear returns slice of the full year, like 85 from 1985.
func SliceYear(year int) string {
	return strconv.Itoa(year)[2:4]
}

// CenturyEqual checks if two given years are of the same century, like 1320 and 1399.
func CenturyEqual(year1 int, year2 int) bool {
	if !IsValidYear(year1) || !IsValidYear(year2) {
		return false
	}
	getCentury := func(year int) string { return strconv.Itoa(year)[0:2] }
	return getCentury(year1) == getCentury(year2)
}

// FormatYearsRangeString formats the range of two years into the string, e.g. "1720–95", or "1720–1805", or "1720–".
// Start year and dash are always present.
// It's supposed to be used for lifespans, meaning we always have birth, but may not have death.
func FormatYearsRangeString(startYear int, finishYear int) string {
	if !IsValidYear(startYear) {
		return ""
	}
	if !IsValidYear(finishYear) {
		return fmt.Sprintf("%d–", startYear)
	}
	if CenturyEqual(startYear, finishYear) {
		return fmt.Sprintf("%d–%s", startYear, SliceYear(finishYear))
	}
	return fmt.Sprintf("%d–%d", startYear, finishYear)
}

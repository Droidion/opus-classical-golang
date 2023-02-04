package utils

import (
	"fmt"
	"strconv"
)

// IsValidYear checks if given string is a 4 digits number, like "1234" (not "-123", "123", or "12345").
func IsValidYear(num int32) bool {
	return num > 999 && num < 10_000
}

// SliceYear returns slice of the full year, like 85 from 1985.
func SliceYear(year int32) string {
	return strconv.FormatInt(int64(year), 10)[2:4]
}

// CenturyEqual checks if two given years are of the same century, like 1320 and 1399.
func CenturyEqual(year1 int32, year2 int32) bool {
	if !IsValidYear(year1) || !IsValidYear(year2) {
		return false
	}
	getCentury := func(year int32) string { return strconv.FormatInt(int64(year), 10)[0:2] }
	return getCentury(year1) == getCentury(year2)
}

// FormatYearsRangeString formats the range of two years into the string, e.g. "1720–95", or "1720–1805", or "1720–".
// Start year and dash are always present.
// It's supposed to be used for lifespans, meaning we always have birth, but may not have death.
func FormatYearsRangeString(startYear int32, finishYear int32) string {
	if !IsValidYear(startYear) && !IsValidYear(finishYear) {
		return ""
	}
	if !IsValidYear(finishYear) {
		return fmt.Sprintf("%d–", startYear)
	}
	if !IsValidYear(startYear) {
		return fmt.Sprintf("%d", finishYear)
	}
	if CenturyEqual(startYear, finishYear) {
		return fmt.Sprintf("%d–%s", startYear, SliceYear(finishYear))
	}
	return fmt.Sprintf("%d–%d", startYear, finishYear)
}

// FormatWorkLength formats minutes into a string with hours and minutes, like "2h 35m"
func FormatWorkLength(lengthInMinutes int32) string {
	hours := lengthInMinutes / 60
	minutes := lengthInMinutes % 60
	if hours == 0 && minutes == 0 {
		return ""
	}
	if hours < 0 || minutes < 0 {
		return ""
	}
	if hours == 0 {
		return fmt.Sprintf("%dm", minutes)
	}
	if minutes == 0 {
		return fmt.Sprintf("%dh", hours)
	}
	return fmt.Sprintf("%dh %dm", hours, minutes)
}

// FormatCatalogueName formats catalogue name of the musical work, like "BWV 12p".
func FormatCatalogueName(catalogueName string, catalogueNumber int32, cataloguePostfix string) string {
	if catalogueName == "" || catalogueNumber == 0 {
		return ""
	}
	return fmt.Sprintf("%s %d%s", catalogueName, catalogueNumber, cataloguePostfix)
}

// FormatWorkName formats music work full name, like "Symphony No. 9 Great".
func FormatWorkName(workTitle string, workNo int32, workNickname string) string {
	if workTitle == "" {
		return ""
	}
	workName := workTitle
	if workNo > 0 {
		workName = workName + fmt.Sprintf(" No. %d", workNo)
	}
	if len(workNickname) > 0 {
		workName = workName + fmt.Sprintf("&nbsp;<em>%s</em>", workNickname)
	}
	return workName
}

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidYearIdentifiesValidYear(t *testing.T) {
	assert.Equal(t, true, IsValidYear(1000))
	assert.Equal(t, true, IsValidYear(1234))
	assert.Equal(t, true, IsValidYear(9999))
}

func TestIsValidYearIdentifiesInvalidYear(t *testing.T) {
	assert.Equal(t, false, IsValidYear(999))
	assert.Equal(t, false, IsValidYear(10000))
	assert.Equal(t, false, IsValidYear(0))
	assert.Equal(t, false, IsValidYear(-1))
}

func TestSliceYearCreatesValidSlice(t *testing.T) {
	assert.Equal(t, "85", SliceYear(1985))
	assert.Equal(t, "99", SliceYear(9999))
}

func TestCenturyEqualReturnsTrueForEqualCenturies(t *testing.T) {
	assert.Equal(t, true, CenturyEqual(1700, 1799))
	assert.Equal(t, true, CenturyEqual(1750, 1749))
}

func TestCenturyEqualReturnsFalseForNonEqualCenturies(t *testing.T) {
	assert.Equal(t, false, CenturyEqual(1699, 1700))
	assert.Equal(t, false, CenturyEqual(1799, 1800))
	assert.Equal(t, false, CenturyEqual(1200, 1500))
	assert.Equal(t, false, CenturyEqual(1, 2))
}

func TestFormatYearsRangeStringFormatsYearsRangeProperly(t *testing.T) {
	assert.Equal(t, "1900–02", FormatYearsRangeString(1900, 1902))
	assert.Equal(t, "1890–1912", FormatYearsRangeString(1890, 1912))
	assert.Equal(t, "1890–", FormatYearsRangeString(1890, 1))
	assert.Equal(t, "1990–", FormatYearsRangeString(1990, 0))
	assert.Equal(t, "1950", FormatYearsRangeString(0, 1950))
	assert.Equal(t, "1912", FormatYearsRangeString(1, 1912))
	assert.Equal(t, "", FormatYearsRangeString(-1, 0))
}

func TestFormatWorkLengthFormatsProperly(t *testing.T) {
	assert.Equal(t, "12m", FormatWorkLength(12))
	assert.Equal(t, "59m", FormatWorkLength(59))
	assert.Equal(t, "1h", FormatWorkLength(60))
	assert.Equal(t, "1h 2m", FormatWorkLength(62))
	assert.Equal(t, "2h 3m", FormatWorkLength(123))
	assert.Equal(t, "", FormatWorkLength(-5))
	assert.Equal(t, "", FormatWorkLength(0))
}

func TestFormatCatalogueName(t *testing.T) {
	assert.Equal(t, "BWV 12m", FormatCatalogueName("BWV", 12, "m"))
	assert.Equal(t, "BWV 12", FormatCatalogueName("BWV", 12, ""))
	assert.Equal(t, "", FormatCatalogueName("", 12, ""))
	assert.Equal(t, "", FormatCatalogueName("BWV", 0, ""))
	assert.Equal(t, "", FormatCatalogueName("", 0, ""))
}

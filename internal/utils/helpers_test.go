package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidYearIdentifiesValidYear(t *testing.T) {
	assert.Equal(t, IsValidYear(1000), true)
	assert.Equal(t, IsValidYear(1234), true)
	assert.Equal(t, IsValidYear(9999), true)
}

func TestIsValidYearIdentifiesInvalidYear(t *testing.T) {
	assert.Equal(t, IsValidYear(999), false)
	assert.Equal(t, IsValidYear(10000), false)
	assert.Equal(t, IsValidYear(0), false)
	assert.Equal(t, IsValidYear(-1), false)
}

func TestSliceYearCreatesValidSlice(t *testing.T) {
	assert.Equal(t, SliceYear(1985), "85")
	assert.Equal(t, SliceYear(9999), "99")
}

func TestCenturyEqualReturnsTrueForEqualCenturies(t *testing.T) {
	assert.Equal(t, CenturyEqual(1700, 1799), true)
	assert.Equal(t, CenturyEqual(1750, 1749), true)
}

func TestCenturyEqualReturnsFalseForNonEqualCenturies(t *testing.T) {
	assert.Equal(t, CenturyEqual(1699, 1700), false)
	assert.Equal(t, CenturyEqual(1799, 1800), false)
	assert.Equal(t, CenturyEqual(1200, 1500), false)
}

func TestFormatYearsRangeStringFormatsYearsRangeProperly(t *testing.T) {
	assert.Equal(t, FormatYearsRangeString(1900, 1902), "1900–02")
	assert.Equal(t, FormatYearsRangeString(1890, 1912), "1890–1912")
	assert.Equal(t, FormatYearsRangeString(1890, 1), "1890–")
	assert.Equal(t, FormatYearsRangeString(1990, 0), "1990–")
	assert.Equal(t, FormatYearsRangeString(1, 1912), "")
	assert.Equal(t, FormatYearsRangeString(-1, 0), "")
}

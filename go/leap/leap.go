package leap

// IsLeapYear reports whether the given year is a leap year.
// The tricky thing here is that a leap year occurs:
// - on every year that is evenly divisible by 4
// - except every year that is evenly divisible by 100
// - unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {
	return isFourthYear(year) && (!isFirstYearOfCentury(year) || isFirstYearOfFourthCentury(year))
}

// isFourthYear reports whether the given year is a fourth year (i.e. divisible by four).
func isFourthYear(year int) bool {
	return year%4 == 0
}

// isFirstYearOfCentury reports whether the given year is the first year of a century (in astronomical year numbering).
func isFirstYearOfCentury(year int) bool {
	return year%100 == 0
}

// isFirstYearOfFourthCentury reports whether the given year is the first year of a fourth century (in astronomical year numbering).
func isFirstYearOfFourthCentury(year int) bool {
	return year%400 == 0
}

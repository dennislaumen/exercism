package leap

// IsLeapYear reports whether the given year is a leap year.
// The tricky thing here is that a leap year occurs:
// - on every year that is evenly divisible by 4
// - except every year that is evenly divisible by 100
// - unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {
	isLeapYear := false

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		isLeapYear = true
	}

	return isLeapYear
}

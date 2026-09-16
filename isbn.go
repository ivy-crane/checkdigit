// Package checkdigit computes and validates the check digits used by book
// and retail product identifiers: ISBN-10, ISBN-13, EAN-13 and UPC-A.
//
// Every exported function is a pure function of its input: no I/O, no
// package-level state, nothing hidden. Validate* functions accept text as
// printed on a book or product (hyphens and spaces allowed); Compute*
// functions take the bare digit prefix and return the check digit that
// belongs after it.
package checkdigit

import "fmt"

// ISBN10CheckDigit computes the check digit for the first 9 digits of an
// ISBN-10. The result is a digit '0'-'9' or 'X' (representing 10).
func ISBN10CheckDigit(digits string) (byte, error) {
	if len(digits) != 9 {
		return 0, fmt.Errorf("checkdigit: ISBN-10 prefix must be 9 digits, got %d", len(digits))
	}
	sum := 0
	for i := 0; i < 9; i++ {
		c := digits[i]
		if !isDigit(c) {
			return 0, fmt.Errorf("checkdigit: non-digit character %q at position %d", c, i+1)
		}
		sum += int(c-'0') * (10 - i)
	}
	check := (11 - (sum % 11)) % 11
	if check == 10 {
		return 'X', nil
	}
	return byte('0' + check), nil
}

// ValidateISBN10 reports whether s is a 10-character ISBN with a correct
// check digit. Hyphens and spaces in s are ignored; a trailing 'x' is
// treated the same as 'X'.
func ValidateISBN10(s string) bool {
	clean := stripSeparators(s)
	if len(clean) != 10 {
		return false
	}
	want, err := ISBN10CheckDigit(clean[:9])
	if err != nil {
		return false
	}
	got := clean[9]
	if got == 'x' {
		got = 'X'
	}
	return got == want
}

// ISBN13CheckDigit computes the check digit for the first 12 digits of an
// ISBN-13 (equivalently, an EAN-13 in the Bookland 978/979 range).
func ISBN13CheckDigit(digits string) (byte, error) {
	if len(digits) != 12 {
		return 0, fmt.Errorf("checkdigit: ISBN-13 prefix must be 12 digits, got %d", len(digits))
	}
	return mod10CheckDigit(digits, 1)
}

// ValidateISBN13 reports whether s is a 13-digit ISBN with a correct check
// digit. Hyphens and spaces in s are ignored.
func ValidateISBN13(s string) bool {
	clean := stripSeparators(s)
	if len(clean) != 13 {
		return false
	}
	want, err := ISBN13CheckDigit(clean[:12])
	if err != nil {
		return false
	}
	return clean[12] == want
}

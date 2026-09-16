package checkdigit

import "fmt"

// EAN13CheckDigit computes the check digit for the first 12 digits of an
// EAN-13 barcode.
func EAN13CheckDigit(digits string) (byte, error) {
	if len(digits) != 12 {
		return 0, fmt.Errorf("checkdigit: EAN-13 prefix must be 12 digits, got %d", len(digits))
	}
	return mod10CheckDigit(digits, 1)
}

// ValidateEAN13 reports whether s is a 13-digit EAN barcode with a correct
// check digit. Hyphens and spaces in s are ignored.
func ValidateEAN13(s string) bool {
	clean := stripSeparators(s)
	if len(clean) != 13 {
		return false
	}
	want, err := EAN13CheckDigit(clean[:12])
	if err != nil {
		return false
	}
	return clean[12] == want
}

// UPCACheckDigit computes the check digit for the first 11 digits of a
// UPC-A barcode.
func UPCACheckDigit(digits string) (byte, error) {
	if len(digits) != 11 {
		return 0, fmt.Errorf("checkdigit: UPC-A prefix must be 11 digits, got %d", len(digits))
	}
	return mod10CheckDigit(digits, 3)
}

// ValidateUPCA reports whether s is a 12-digit UPC-A barcode with a correct
// check digit. Hyphens and spaces in s are ignored.
func ValidateUPCA(s string) bool {
	clean := stripSeparators(s)
	if len(clean) != 12 {
		return false
	}
	want, err := UPCACheckDigit(clean[:11])
	if err != nil {
		return false
	}
	return clean[11] == want
}

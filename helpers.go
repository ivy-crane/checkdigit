package checkdigit

import "fmt"

// stripSeparators removes the hyphens and spaces publishers and GS1 use to
// group digits for readability. Callers of Validate* pass raw label text;
// Compute* functions expect already-clean digit strings.
func stripSeparators(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '-' || c == ' ' {
			continue
		}
		out = append(out, c)
	}
	return string(out)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// mod10CheckDigit implements the weighted mod-10 scheme shared by EAN-13,
// UPC-A and ISBN-13: weights alternate between firstWeight and its
// complement (4 - firstWeight, so 1<->3), starting from the leftmost digit.
func mod10CheckDigit(digits string, firstWeight int) (byte, error) {
	secondWeight := 4 - firstWeight
	sum := 0
	weight := firstWeight
	for i := 0; i < len(digits); i++ {
		c := digits[i]
		if !isDigit(c) {
			return 0, fmt.Errorf("checkdigit: non-digit character %q at position %d", c, i+1)
		}
		sum += int(c-'0') * weight
		if weight == firstWeight {
			weight = secondWeight
		} else {
			weight = firstWeight
		}
	}
	check := (10 - (sum % 10)) % 10
	return byte('0' + check), nil
}

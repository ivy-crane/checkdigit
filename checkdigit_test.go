package checkdigit

import "testing"

func TestISBN10CheckDigit(t *testing.T) {
	cases := []struct {
		name    string
		prefix  string
		want    byte
		wantErr bool
	}{
		{name: "all ones", prefix: "111111111", want: '1'},
		{name: "wikipedia example", prefix: "030640615", want: '2'},
		{name: "check digit is X", prefix: "080442957", want: 'X'},
		{name: "too short", prefix: "12345678", wantErr: true},
		{name: "too long", prefix: "1234567890", wantErr: true},
		{name: "non-digit", prefix: "01234567a", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ISBN10CheckDigit(c.prefix)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ISBN10CheckDigit(%q) = %c, nil; want error", c.prefix, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ISBN10CheckDigit(%q) returned error: %v", c.prefix, err)
			}
			if got != c.want {
				t.Errorf("ISBN10CheckDigit(%q) = %c, want %c", c.prefix, got, c.want)
			}
		})
	}
}

func TestValidateISBN10(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "hyphenated", in: "0-306-40615-2", want: true},
		{name: "no separators", in: "0306406152", want: true},
		{name: "spaced", in: "0 306 40615 2", want: true},
		{name: "check digit X", in: "0-8044-2957-X", want: true},
		{name: "lowercase x accepted", in: "0-8044-2957-x", want: true},
		{name: "wrong check digit", in: "0-306-40615-3", want: false},
		{name: "wrong length", in: "0-306-4061-2", want: false},
		{name: "non-digit body", in: "0-30A-40615-2", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ValidateISBN10(c.in); got != c.want {
				t.Errorf("ValidateISBN10(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestISBN13CheckDigit(t *testing.T) {
	cases := []struct {
		name    string
		prefix  string
		want    byte
		wantErr bool
	}{
		{name: "wikipedia example", prefix: "978013419044", want: '0'},
		{name: "979 range", prefix: "979102345678", want: '3'},
		{name: "too short", prefix: "97801341904", wantErr: true},
		{name: "non-digit", prefix: "97801341904a", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ISBN13CheckDigit(c.prefix)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ISBN13CheckDigit(%q) = %c, nil; want error", c.prefix, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ISBN13CheckDigit(%q) returned error: %v", c.prefix, err)
			}
			if got != c.want {
				t.Errorf("ISBN13CheckDigit(%q) = %c, want %c", c.prefix, got, c.want)
			}
		})
	}
}

func TestValidateISBN13(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "hyphenated", in: "978-0-13-419044-0", want: true},
		{name: "no separators", in: "9780134190440", want: true},
		{name: "wrong check digit", in: "978-0-13-419044-1", want: false},
		{name: "wrong length", in: "978-0-13-41904-0", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ValidateISBN13(c.in); got != c.want {
				t.Errorf("ValidateISBN13(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestISBN10ToISBN13(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "plain", in: "0306406152", want: "9780306406157"},
		{name: "hyphenated", in: "0-13-419044-0", want: "9780134190440"},
		{name: "check digit X", in: "0-8044-2957-X", want: "9780804429573"},
		{name: "bad check digit", in: "0-306-40615-3", wantErr: true},
		{name: "empty", in: "", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ISBN10ToISBN13(c.in)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ISBN10ToISBN13(%q) = %q, nil; want error", c.in, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ISBN10ToISBN13(%q) returned error: %v", c.in, err)
			}
			if got != c.want {
				t.Errorf("ISBN10ToISBN13(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestISBN13ToISBN10(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "plain", in: "9780306406157", want: "0306406152"},
		{name: "hyphenated", in: "978-0-13-419044-0", want: "0134190440"},
		{name: "check digit X", in: "9780804429573", want: "080442957X"},
		{name: "979 has no ISBN-10", in: "9791023456783", wantErr: true},
		{name: "bad check digit", in: "9780306406150", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ISBN13ToISBN10(c.in)
			if c.wantErr {
				if err == nil {
					t.Fatalf("ISBN13ToISBN10(%q) = %q, nil; want error", c.in, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ISBN13ToISBN10(%q) returned error: %v", c.in, err)
			}
			if got != c.want {
				t.Errorf("ISBN13ToISBN10(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestEAN13CheckDigit(t *testing.T) {
	cases := []struct {
		name    string
		prefix  string
		want    byte
		wantErr bool
	}{
		{name: "known good", prefix: "590123412345", want: '7'},
		{name: "too long", prefix: "5901234123457", wantErr: true},
		{name: "non-digit", prefix: "59012341234a", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := EAN13CheckDigit(c.prefix)
			if c.wantErr {
				if err == nil {
					t.Fatalf("EAN13CheckDigit(%q) = %c, nil; want error", c.prefix, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("EAN13CheckDigit(%q) returned error: %v", c.prefix, err)
			}
			if got != c.want {
				t.Errorf("EAN13CheckDigit(%q) = %c, want %c", c.prefix, got, c.want)
			}
		})
	}
}

func TestValidateEAN13(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "known good", in: "5901234123457", want: true},
		{name: "hyphenated", in: "590-123412-3457", want: true},
		{name: "wrong check digit", in: "5901234123450", want: false},
		{name: "wrong length", in: "590123412345", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ValidateEAN13(c.in); got != c.want {
				t.Errorf("ValidateEAN13(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestUPCACheckDigit(t *testing.T) {
	cases := []struct {
		name    string
		prefix  string
		want    byte
		wantErr bool
	}{
		{name: "known good", prefix: "04900002891", want: '1'},
		{name: "too short", prefix: "0490000289", wantErr: true},
		{name: "non-digit", prefix: "0490000289a", wantErr: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := UPCACheckDigit(c.prefix)
			if c.wantErr {
				if err == nil {
					t.Fatalf("UPCACheckDigit(%q) = %c, nil; want error", c.prefix, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("UPCACheckDigit(%q) returned error: %v", c.prefix, err)
			}
			if got != c.want {
				t.Errorf("UPCACheckDigit(%q) = %c, want %c", c.prefix, got, c.want)
			}
		})
	}
}

func TestValidateUPCA(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "known good", in: "049000028911", want: true},
		{name: "hyphenated", in: "0-49000-02891-1", want: true},
		{name: "wrong check digit", in: "049000028912", want: false},
		{name: "wrong length", in: "04900002891", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ValidateUPCA(c.in); got != c.want {
				t.Errorf("ValidateUPCA(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

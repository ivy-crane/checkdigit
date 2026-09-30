# checkdigit

Check digit computation and validation for ISBN-10, ISBN-13, EAN-13 and
UPC-A identifiers.

Every one of these formats ends in a digit that's derived from the digits
before it, by a different weighting scheme per format. Publishers and
retailers use it to catch typos and scanner misreads before they turn into
a wrong book or a wrong price. This package implements the four schemes
that show up most often: ISBN-10's mod-11, and the mod-10 scheme shared
(with different starting weights) by ISBN-13, EAN-13 and UPC-A.

## Why a separate function per format

The four formats look similar but aren't interchangeable: ISBN-10 uses
mod 11 with weights 10 down to 1 and can produce an 'X' check digit, while
ISBN-13/EAN-13/UPC-A use mod 10 with alternating weights of 1 and 3 (which
digit gets weight 3 first differs between UPC-A and the other two). Rather
than one function that branches on a format argument, each format gets its
own `Compute*`/`Validate*` pair, so the type signature tells you what
you're allowed to pass in.

## Install

```
go get github.com/ivy-crane/checkdigit
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/ivy-crane/checkdigit"
)

func main() {
	// Validate a full identifier as printed on the item.
	fmt.Println(checkdigit.ValidateISBN13("978-0-13-419044-0")) // true
	fmt.Println(checkdigit.ValidateISBN10("0-13-419044-0"))     // true
	fmt.Println(checkdigit.ValidateEAN13("5901234123457"))      // true
	fmt.Println(checkdigit.ValidateUPCA("049000028911"))        // true

	// Compute the check digit for a prefix you already have, e.g. when
	// assigning a new identifier.
	check, err := checkdigit.ISBN13CheckDigit("978013419044")
	if err != nil {
		panic(err)
	}
	fmt.Printf("check digit: %c\n", check) // check digit: 0
}
```

Every `Compute*` function takes the bare digit prefix (no hyphens) and
returns the single check digit that belongs after it. Every `Validate*`
function takes the full identifier as it would appear on a book cover or
package, hyphens and spaces included, and reports whether the trailing
check digit is correct.

None of these functions do I/O or keep state, so they compose directly
into whatever validation or generation code calls them.

## Status

Early. The four formats above are implemented and tested against known
reference numbers. `ISBN10ToISBN13` and `ISBN13ToISBN10` convert between
the two ISBN forms (only 978-prefixed ISBN-13s have an ISBN-10 form). Not
yet covered: EAN-8 and a batch-checking helper for reading a list of
identifiers from a file.

## License

MIT, see [LICENSE](LICENSE).

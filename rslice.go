/*
package rslice provides some common []rune patterns for moving runes around
within a slice, justifying embedded text, etc

Charles <asciifaceman> Corbett 2023

MIT License
*/
package rslice

import (
	"fmt"
	"unicode"
)

// Whitespace returns true if the entire []rune is whitespace
// It will also return true if the slice is empty
func Whitespace(slice []rune) bool {
	for _, char := range slice {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// Words returns a count of non-whitespace groupings of characters
// that may or may not be a word
//
// Can be useful to discover how many areas of whitespace you have for
// purposes such as full text justification across a width
//
// Does not recognize control characters (such as \n or \t) as non-whitespace
// characters as per unicode stdlib
func Words(slice []rune) int {
	count := 0
	word := false

	for _, char := range slice {
		if !unicode.IsSpace(char) {
			if !word {
				word = true
				count++
			}
		} else {
			if word {
				word = false
			}
		}
	}
	return count

}

// Valid returns true a slice has width and is not all whitespace
func Valid(slice []rune) bool {
	if len(slice) > 0 && !Whitespace(slice) {
		return true
	}

	return false
}

// ShiftLeft shifts the rune slice one to the left and returns a copy
// if the slice is not all whitespace
func ShiftLeft(slice []rune) []rune {
	if !Valid(slice) {
		return slice
	}

	return append(slice[1:], slice[0])
}

// ShiftRight shits the rune slice one to the right and returns a copy
// if the slice is not all whitespace
func ShiftRight(slice []rune) []rune {
	if !Valid(slice) {
		return slice
	}

	return append(slice[len(slice)-1:], slice[:len(slice)-1]...)
}

// ShiftWhitespaceLeft shifts any whitespace right of the last non-whitespace
// character to the left of the first non-whitespace character in the rune slice
// and returns a copy if the slice is not all whitespace
func ShiftWhitespaceLeft(slice []rune) []rune {
	if !Valid(slice) {
		return slice
	}

	if unicode.IsSpace(slice[len(slice)-1]) {
		slice = ShiftRight(slice)
		return ShiftWhitespaceLeft(slice)
	} else {
		return slice
	}
}

// ShiftWhitespaceRight shifts any whitespace left of the first non-whitespace
// character to the right of the last non-whitespace character in the rune sluce
// and returns a copy if the slice is not all whitespace
func ShiftWhitespaceRight(slice []rune) []rune {
	if !Valid(slice) {
		return slice
	}

	if unicode.IsSpace(slice[0]) {
		slice = ShiftLeft(slice)
		return ShiftWhitespaceRight(slice)
	} else {
		return slice
	}
}

// Newline returns true if the given rune is a Linux, Darwin, or Windows newline character
func Newline(r rune) bool {
	if unicode.IsControl(r) {
		if r == rune('\r') || r == rune('\n') {
			return true
		}
	}
	return false
}

// TrimExcessWhitespace will remove any occurance of whitespace greater
// than one index.
func TrimExcessWhitespace(slice []rune) []rune {
	count := 0
	for _, r := range slice {
		if unicode.IsSpace(r) {
			count++
		} else {
			count = 0
		}

	}

	return slice
}

// LeastWhitespaceIndex returns an index point of a []rune with
// the least whitespace between the left and right
// most characters
//
// It will wait until it has passed at least one non-whitespace character
// before recording the potential index.
//
// A return of -1 indicates there is no suitable index, an example
// string would be ` a ` which has no whitespace between two non-whitespace
// characters
//
// Currently this function will trigger an ignore on whitespace after a control
// character is encountered until the next non-whitespace non-control character
// and effectively erase any whitespace between the previous non-ws/non-cc rune
// to prevent returning an index between or before a control character wh
func LeastWhitespaceIndex(slice []rune) int {
	var idx int
	count := len(slice)
	word := false
	ignore := false

	subcount := count
	for i, r := range slice {
		if i == len(slice) {
			break
		}
		if unicode.IsControl(r) {
			// ignore any new whitespaces until the
			// next non-whitespace character
			ignore = true
			continue
		}
		if !unicode.IsSpace(r) {
			if !word {
				word = true
				if !ignore {
					// count the index if not ignored
					if subcount < count {
						idx = i
						count = subcount
					}
					subcount = 0
				} else {
					// if ignored reset the count to terminate
					// any whitespace since the last valid word char
					subcount = 0
				}
			}
			// disable ignore if it is a word boundary
			ignore = false
			subcount = 0
		} else {
			if word {
				word = false
			}
			// increase count of this segment of whitespace
			subcount++
		}
	}
	return idx - 1
}

/*
NormalizeWhitespace takes the left and right whitespace of the given
rune slice and spreads it across the interior whitespace of the rune
slice between the inner and outer most non-whitespace character

Maintains the []rune's width

Returns the slice unchanged if the []rune contains only whitespace,
has no length, or has fewer than 2 words since there would be no
inner whitespace to utilize

Usage:
```go

	s := []rune("    A string with whitespace to the left and right    ")
	s = rslice.NormalizeWhitespace(s)

	// s should now be "A  string  with  whitespace  to  the   left and  right"

```
*/
func NormalizeWhitespace(slice []rune) []rune {
	wordCount := Words(slice)

	if Whitespace(slice) || len(slice) < 1 || wordCount < 2 {
		return slice
	}

	slice = ShiftWhitespaceLeft(slice)
	slice = Normalize(slice)

	return slice
}

/*
Normalize is a recursive function that will take all whitespace left
of the left most non-whitespace and non-control-character space and move it
somewhere in the interior starting on the left most interior and working in

Normalize maintains the []rune's width

Example
```go

	s := []rune("    A string with whitespace to the left")
	s = rslice.Normalize(s)

	// s should now be "A  string  with  whitespace  to the left"

```
*/
func Normalize(slice []rune) []rune {
	if !Valid(slice) {
		return slice
	}

	if unicode.IsSpace(slice[0]) {
		d := LeastWhitespaceIndex(slice)
		if d == -1 {
			return slice
		}
		fmt.Println(d)
		slice = slice[1:]
		slice = append(slice[:d+1], slice[d:]...)
		slice[d] = rune(' ')
		return Normalize(slice)
	} else {
		return slice
	}
}

/*
package rslice provides some common []rune patterns for moving runes around
within a slice, justifying embedded text, etc

Charles <asciifaceman> Corbett 2023

MIT License
*/
package rslice

import "unicode"

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

//

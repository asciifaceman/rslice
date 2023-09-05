package rslice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestWhitespace(t *testing.T) {
	tests := map[string]bool{
		"":                      true, // this is an odd one and I don't know
		" ":                     true,
		"     ":                 true,
		strings.Repeat(" ", 50): true,
		"t":                     false,
		"  t  ":                 false,
		fmt.Sprintf("%sTHING", strings.Repeat(" ", 50)): false,
	}

	for tc := range tests {
		slice := []rune(tc)

		if w := Whitespace(slice); w != tests[tc] {
			t.Fatalf("Expected Whitespace result of [%t] for string (%s) but got [%t]", tests[tc], tc, w)
		}
	}
}

func TestWords(t *testing.T) {
	tests := map[string]int{
		"  abc  def      ghi":      3,
		" a c g y 34           35": 6,
		" abc def  _   feff  ":     4,
		"\n f23 \n \t":             1,
	}

	for tc := range tests {
		slice := []rune(tc)

		if w := Words(slice); w != tests[tc] {
			t.Fatalf("Expected string [%s] to have [%d] 'words' but counted [%d]", tc, tests[tc], w)
		}
	}
}

func TestShiftLeft(t *testing.T) {
	tests := map[string][]rune{
		"":          []rune(""),
		" ":         []rune(" "),
		"     ":     []rune("     "),
		"ABCDEF":    []rune("BCDEFA"),
		" gh%Y^uio": []rune("gh%Y^uio "),
	}

	for tc := range tests {
		slice := []rune(tc)

		//t.Fatalf("Expected given string to be rune-shifted one to the left\nEXPECTED: [%s][%v]\nRECEIVED: [%s][%v]", )

		if s := ShiftLeft(slice); !reflect.DeepEqual(s, tests[tc]) {
			t.Fatalf("Expected given string to be rune-shifted one to the left\nEXPECTED: [%s][%v]\nRECEIVED: [%s][%v]\n", string(tests[tc]), tests[tc], string(s), s)
		}

	}
}

func TestShiftRight(t *testing.T) {
	tests := map[string][]rune{
		"":          []rune(""),
		" ":         []rune(" "),
		"     ":     []rune("     "),
		"ABCDEF":    []rune("FABCDE"),
		" gh%Y^uio": []rune("o gh%Y^ui"),
	}

	for tc := range tests {
		slice := []rune(tc)

		if s := ShiftRight(slice); !reflect.DeepEqual(s, tests[tc]) {
			t.Fatalf("Expected given string to be rune-shifted one to the left\nEXPECTED: [%s][%v]\nRECEIVED: [%s][%v]\n", string(tests[tc]), tests[tc], string(s), s)
		}

	}
}

func TestShiftWhiteSpaceLeft(t *testing.T) {
	tests := map[string][]rune{
		"":              []rune(""),
		" ":             []rune(" "),
		"     ":         []rune("     "),
		"ABCDEF ":       []rune(" ABCDEF"),
		"ABC DEF    ":   []rune("    ABC DEF"),
		"  ABC DEF    ": []rune("      ABC DEF"),
	}

	for tc := range tests {
		slice := []rune(tc)

		if s := ShiftWhitespaceLeft(slice); !reflect.DeepEqual(s, tests[tc]) {
			t.Fatalf("Expected given string to have white space on the right rune-shifted to the left\nEXPECTED: [%s][%v]\nRECEIVED:[%s][%v]", string(tests[tc]), tests[tc], string(s), s)
		}
	}
}

func TestShiftWhiteSpaceRight(t *testing.T) {
	tests := map[string][]rune{
		"":                []rune(""),
		" ":               []rune(" "),
		"     ":           []rune("     "),
		" ABCDEF":         []rune("ABCDEF "),
		"    ABC DEF    ": []rune("ABC DEF        "),
		"    ABC DEF  ":   []rune("ABC DEF      "),
	}

	for tc := range tests {
		slice := []rune(tc)

		if s := ShiftWhitespaceRight(slice); !reflect.DeepEqual(s, tests[tc]) {
			t.Fatalf("Expected given string to have white space on the right rune-shifted to the left\nEXPECTED: [%s][%v]\nRECEIVED:[%s][%v]", string(tests[tc]), tests[tc], string(s), s)
		}
	}
}

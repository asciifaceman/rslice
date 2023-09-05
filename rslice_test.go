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

func TestNewline(t *testing.T) {
	tests := map[string]bool{
		"\n": true,
		"\a": false,
		"\r": true,
	}

	for tc := range tests {
		r := []rune(tc)
		if b := Newline(r[0]); b != tests[tc] {
			t.Fatalf("Expected character [%v][%v] to newline [%t] but got [%t]", tc, r[0], tests[tc], b)
		}
	}
}

func TestTrimExcessWhitespace(t *testing.T) {
	tests := map[string][]rune{
		"This  string  has  excess  whitespace": []rune("This string has excess whitespace"),
		"   This string":                        []rune(" This string"),
	}

	for tc := range tests {
		new := TrimExcessWhitespace([]rune(tc))
		if !reflect.DeepEqual(new, tests[tc]) {
			t.Fatalf("derp")
		}

	}
}

func TestLeastWhitespaceIndex(t *testing.T) {
	tests := map[string][]int{
		" A  test  to  make  sure   the  right  position  is  chosen": {
			3, 10,
		},
		"    ": {
			-1, -1,
		},
		"   Another test to make sure it doesn't choose the      zero 0:2 index": {
			10, 16,
		},
		"\tAnother test \n\n    hur de\tdurrr\n  ": {
			8, 24,
		},
		" a ": {
			-1, -1,
		},
		" a b": {
			2, 3,
		},
		" a   ": {
			-1, -1,
		},
		"\n\t\t\a": {
			-1, -1,
		},
		" a \n \r\n  \n \t \t": {
			-1, -1,
		},
		" a\n\t \r\tb c": {
			8, 9,
			// TODO: this test case indicates a problem with unbalanced growth
			// it will always pick between b and c
		},
		"a \nb c": {
			4, 5,
			// TODO: this test case indicates a problem with unbalanced growth
			// it will always pick between b and c
		},
		"A  test\nwith  more\r\ncontrol  characters  in  compromising  positions": {
			2, 14,
		},
		" A  Test\nwith  more\r\ncontrol  characters again": {
			40, 3,
		},
		"A   tricky\n string test": {
			18, 19,
		},
		"A  tricky  test  with\n \n \ncontrol characters  in  potentially  breaking  positions": {
			33, 2,
		},
		"Another   tricky   test \n \n \nwhere  control  characters  mess  with  things": {
			35, 45,
		},
	}

	/*
		I am testing with a lot of control characters for stability despite the fact
		that I will likely have processed & stripped them already when using
		LeastWhitespaceIndex to do final justification/formatting

		I am torn on whether to place the index before or after control characters
		or treat them like a whitespace
	*/

	for tc := range tests {
		rs := []rune(tc)
		d := LeastWhitespaceIndex(rs)
		if d != tests[tc][0] {
			t.Fatalf("Expected string [%s] to have least whitespace at index [%d] but function returned [%d]", tc, tests[tc][0], d)
		}

		if d == -1 {
			continue
		}

		t.Log(string(rs))
		// find runner up
		rs = append(rs[:d+1], rs[d:]...)
		t.Log(string(rs))
		if d2 := LeastWhitespaceIndex(rs); d2 != tests[tc][1] {
			t.Fatalf("Runner up expected [%v] to produce index [%d] but got [%d]  (testmap: %d:%d)", string(rs), tests[tc][1], d2, tests[tc][0], tests[tc][1])
		}
	}

}

func TestNormalizeWhitespace(t *testing.T) {
	tests := map[string][]rune{
		"":                            []rune(""),
		"     ":                       []rune("     "),
		"word":                        []rune("word"),
		"    Test something here":     []rune("Test   something   here"),
		"   Test something here":      []rune("Test   something  here"),
		"    Test Something Here    ": []rune("Test     Something     Here"),
		"     Another   tricky   test \n \n \nwhere  control  characters  mess  with  things": []rune("Another   tricky   test \n \n \nwhere   control   characters   mess   with   things"),
		"   a\n\t \r\tb c": []rune("a\n\t \r\tb    c"), // not great but expected
	}

	for tc := range tests {
		rs := []rune(tc)
		width := len(rs)
		rs = NormalizeWhitespace(rs)

		t.Logf("Original string [w: %d][%s]", width, tc)
		t.Logf("New String      [w: %d][%s]", len(rs), string(rs))
		for i := 0; i < len(tests[tc]); i++ {
			if rs[i] != tests[tc][i] {
				t.Logf("Expected: [%v]\nReceived: [%v]\n", []rune(tests[tc]), rs)
				t.Fatalf("Expected normalized string to appear like [w: %d][%s]\nbut got [w: %d][%s]", len(tests[tc]), string(tests[tc]), len(rs), string(rs))
			}
		}

		//		if reflect.DeepEqual(rs, tests[tc]) {
		//			t.Logf("Expected: [%v]\nReceived: [%v]\n", []rune(tests[tc]), rs)
		//			t.Fatalf("Expected normalized string to appear like [w: %d][%s]\nbut got [w: %d][%s]", len(tests[tc]), string(tests[tc]), len(rs), string(rs))
		//		}
	}
}

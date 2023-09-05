# rslice

![test workflow](https://github.com/asciifaceman/rslice/actions/workflows/main-test.yml/badge.svg) [![Go Coverage](https://github.com/asciifaceman/rslice/wiki/coverage.svg)](https://raw.githack.com/wiki/asciifaceman/rslice/coverage.html) [![Go Report Card](https://goreportcard.com/badge/github.com/asciifaceman/rslice)](https://goreportcard.com/report/github.com/asciifaceman/rslice) [![Go Reference](https://pkg.go.dev/badge/github.com/asciifaceman/rslice.svg)](https://pkg.go.dev/github.com/asciifaceman/rslice)


Some operations for working with `[]rune` thingies and stuff which may be particularly
useful for things like justifying text, manipulating strings, or just being a dang hunk.

I mostly wrote this to lift it out of [tooey](github.com/asciifaceman/tooey)

# Some Example Usage

```go

s := []rune("    TEST STRING")
s = ShiftWhiteSpaceRight(s)

> "TEST STRING     "

s := []rune("ABCDEF")
s = ShiftLeft(s)

> "BCDEFA"

```
/*
Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like unicode.IsLetter.

Charcount computes counts of Unicode characters
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func unicodeType(r rune) string {
	switch {
	case unicode.IsControl(r):
		return "control"
	case unicode.IsDigit(r):
		return "digit"
	case unicode.IsLetter(r):
		return "letter"
	case unicode.IsLower(r):
		return "lower"
	case unicode.IsMark(r):
		return "mark"
	case unicode.IsNumber(r):
		return "number"
	case unicode.IsPunct(r):
		return "punct"
	default:
		return "unknown"
	}
}

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	invalid := 0                   // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[unicodeType(r)]++
	}
	fmt.Printf("unicode category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
